/*
Copyright 2017 The Kubernetes Authors.

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

package usercluster

import (
	"fmt"
	"os"

	"github.com/golang/glog"

	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/rest"
	fedv1 "k8s.io/federation/apis/federation/v1beta1"
	"k8s.io/federation/pkg/federation-controller/util/identityprovider"
)

type UserClusterIdentityProvider struct {
	lister *userClusterIdentityLister
}

func NewUserClusterIdentityProvider(apiextClientConfig *rest.Config, identityNamespace string) (*UserClusterIdentityProvider, error) {
	apiextClient, err := apiextensionsclient.NewForConfig(apiextClientConfig)
	if err != nil {
		return nil, fmt.Errorf("failed creating apiextension client, err: %v", err)
	}

	_, err = CreateCrd(apiextClient)
	if err != nil {
		return nil, fmt.Errorf("failed creating UserClusterIdentity CRD, err: %v", err)
	}

	client, _, err := NewCRDClientForConfig(apiextClientConfig)
	if err != nil {
		return nil, fmt.Errorf("failed creating UserClusterIdentity CRD client, err: %v", err)
	}

	return &UserClusterIdentityProvider{
		lister: NewUserClusterIdentityLister(client.UserClusterIdentity(identityNamespace)),
	}, nil
}

func NewInClusterUserClusterIdentityProviderOrDie() *UserClusterIdentityProvider {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(fmt.Errorf("error in creating in-cluster config: %s", err))
	}

	namespace := os.Getenv("POD_NAMESPACE")
	if namespace == "" {
		panic(fmt.Errorf("unexpected: POD_NAMESPACE env var returned empty string"))
	}

	provider, err := NewUserClusterIdentityProvider(config, namespace)
	if err != nil {
		panic(fmt.Errorf("error in creating identity provider: %s", err))
	}

	return provider
}

func (p *UserClusterIdentityProvider) GetUserIdentityForCluster(username string, cluster *fedv1.Cluster) (*identityprovider.Identity, error) {
	identities, err := p.lister.ListByUserCluster(username, cluster.Name)
	if err != nil {
		return nil, fmt.Errorf("error listing UserClusterIdentity objects, err: %v", err)
	}

	if len(identities) == 0 {
		return nil, fmt.Errorf("identity not found for user %v on cluster %v", username, cluster.Name)
	}

	// TODO: handle duplications
	// TODO: prevent mutating Spec.Username and Spec.ClusterName, OpenAPI v3 CRD validation
	if len(identities) > 1 {
		glog.Warningf("multiple (%v) identities found for user %v on cluster %v", len(identities), username, cluster.Name)
	}
	return &identities[0].Spec.Identity, nil
}
