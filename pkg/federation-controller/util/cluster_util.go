/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/golang/glog"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilnet "k8s.io/apimachinery/pkg/util/net"
	"k8s.io/apimachinery/pkg/util/wait"
	clientsetv1 "k8s.io/client-go/kubernetes"
	clientsetcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	federation_v1beta1 "k8s.io/federation/apis/federation/v1beta1"
	"k8s.io/kubernetes/pkg/serviceaccount"
)

const (
	KubeAPIQPS              = 20.0
	KubeAPIBurst            = 30
	KubeconfigSecretDataKey = "kubeconfig"
	getSecretTimeout        = 1 * time.Minute
)

func findServerAddressFromCidrs(addByCidrs []federation_v1beta1.ServerAddressByClientCIDR) (string, error) {
	hostIP, err := utilnet.ChooseHostInterface()
	if err != nil {
		return "", err
	}
	for _, item := range addByCidrs {
		_, cidrnet, err := net.ParseCIDR(item.ClientCIDR)
		if err != nil {
			return "", err
		}
		myaddr := net.ParseIP(hostIP.String())
		if cidrnet.Contains(myaddr) == true {
			return item.ServerAddress, nil
		}
	}
	return "", fmt.Errorf("unable to find server address for client IP %v", hostIP)
}

func BuildClusterConfig(c *federation_v1beta1.Cluster) (*restclient.Config, error) {
	serverAddress, err := findServerAddressFromCidrs(c.Spec.ServerAddressByClientCIDRs)
	if err != nil {
		return nil, err
	}

	var clusterConfig *restclient.Config
	if serverAddress != "" {
		if c.Spec.SecretRef == nil {
			glog.Infof("didn't find secretRef for cluster %s. Trying insecure access", c.Name)
			clusterConfig, err = clientcmd.BuildConfigFromFlags(serverAddress, "")
		} else {
			if c.Spec.SecretRef.Name == "" {
				return nil, fmt.Errorf("found secretRef but no secret name for cluster %s", c.Name)
			}
			secret, err := getSecret(c.Spec.SecretRef.Name)
			if err != nil {
				return nil, err
			}
			clusterConfig, err = buildConfigFromSecret(secret, serverAddress)
		}
		if err != nil {
			return nil, err
		}
		clusterConfig.QPS = KubeAPIQPS
		clusterConfig.Burst = KubeAPIBurst
	}
	return clusterConfig, nil
}

func BuildClusterConfigFromRef(corev1Client clientsetcorev1.CoreV1Interface, cluster *federation_v1beta1.Cluster, credRef corev1.ObjectReference) (*restclient.Config, error) {

	serverAddress, err := findServerAddressFromCidrs(cluster.Spec.ServerAddressByClientCIDRs)
	if err != nil {
		return nil, err
	}

	if credRef.Kind != "ServiceAccount" && credRef.Kind != "Secret" {
		return nil, fmt.Errorf("cluster credential should be either ServiceAccount or Secret. Actual Kind: %v, namespace/name: %v/%v",
			credRef.Kind, credRef.Namespace, credRef.Name)
	}

	var configSecret *corev1.Secret
	if credRef.Kind == "ServiceAccount" {
		sa, err := corev1Client.ServiceAccounts(credRef.Namespace).Get(credRef.Name, metav1.GetOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed get service account %v/%v, err: %v", credRef.Namespace, credRef.Name, err)
		}

		// find the correct secret for the service account
		getOptions := metav1.GetOptions{}
		for _, secretRef := range sa.Secrets {
			secret, err := corev1Client.Secrets(sa.Namespace).Get(secretRef.Name, getOptions)
			if err == nil && serviceaccount.IsServiceAccountToken(secret, sa) {
				configSecret = secret
				break
			}
		}
	} else /*credRef.Kind == "Secret"*/ {
		secret, err := corev1Client.Secrets(credRef.Namespace).Get(credRef.Name, metav1.GetOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed get secret %v/%v, err: %v", credRef.Namespace, credRef.Name, err)
		}
		configSecret = secret
	}

	return buildConfigFromSecret(configSecret, serverAddress)
}

func buildConfigFromSecret(secret *corev1.Secret, serverAddress string) (*restclient.Config, error) {
	var clusterConfig *restclient.Config
	var err error
	// Pre-1.7, the secret contained a serialized kubeconfig which contained appropriate credentials.
	// Post-1.7, the secret contains credentials for a service account.
	// Check for the service account credentials, and use them if they exist; if not, use the
	// serialized kubeconfig.
	token, tokenFound := secret.Data["token"]
	ca, caFound := secret.Data["ca.crt"]
	if tokenFound != caFound {
		return nil, fmt.Errorf("secret should have values for either both 'ca.crt' and 'token' in its Data, or neither: %v", secret)
	} else if tokenFound && caFound {
		clusterConfig, err = clientcmd.BuildConfigFromFlags(serverAddress, "")
		clusterConfig.CAData = ca
		clusterConfig.BearerToken = string(token)
	} else {
		clusterConfig, err = clientcmd.BuildConfigFromKubeconfigGetter(serverAddress, KubeconfigGetterForSecret(secret))
	}

	if err != nil {
		return nil, err
	}

	clusterConfig.QPS = KubeAPIQPS
	clusterConfig.Burst = KubeAPIBurst

	return clusterConfig, nil
}

func GetInClusterClient() (clientsetv1.Interface, error) {
	cc, err := restclient.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("error in creating in-cluster config: %s", err)
	}
	return clientsetv1.NewForConfig(cc)
}

// getSecret gets a secret from the cluster.
func getSecret(secretName string) (*corev1.Secret, error) {
	// Get the namespace this is running in from the env variable.
	namespace := os.Getenv("POD_NAMESPACE")
	if namespace == "" {
		return nil, fmt.Errorf("unexpected: POD_NAMESPACE env var returned empty string")
	}
	// Get a client to talk to the k8s apiserver, to fetch secrets from it.
	client, err := GetInClusterClient()
	if err != nil {
		return nil, fmt.Errorf("error in creating in-cluster client: %s", err)
	}

	var secret *corev1.Secret
	err = wait.PollImmediate(1*time.Second, getSecretTimeout, func() (bool, error) {
		secret, err = client.CoreV1().Secrets(namespace).Get(secretName, metav1.GetOptions{})
		if err == nil {
			return true, nil
		}
		glog.Warningf("error in fetching secret: %s", err)
		return false, nil
	})
	if err != nil {
		return nil, fmt.Errorf("timed out waiting for secret: %s", err)
	}
	if secret == nil {
		return nil, fmt.Errorf("unexpected: received null secret %s", secretName)
	}
	return secret, nil
}

// KubeconfigGetterForSecret gets the kubeconfig from the given secret.
// This is to inject a different KubeconfigGetter in tests. We don't use
// the standard one which calls NewInCluster in tests to avoid having to
// set up service accounts and mount files with secret tokens.
var KubeconfigGetterForSecret = func(secret *corev1.Secret) clientcmd.KubeconfigGetter {
	return func() (*clientcmdapi.Config, error) {
		data, ok := secret.Data[KubeconfigSecretDataKey]
		if !ok {
			return nil, fmt.Errorf("secret does not have data with key %s", KubeconfigSecretDataKey)
		}
		return clientcmd.Load(data)
	}
}
