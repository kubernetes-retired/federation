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

package single

import (
	"io"
	"os"

	corev1 "k8s.io/api/core/v1"
	fedv1 "k8s.io/federation/apis/federation/v1beta1"
	"k8s.io/federation/pkg/federation-controller/util/identityprovider"
)

const (
	providerName = "single"
)

type singleIdentityProvider struct{}

func NewSingleIdentityProvider() *singleIdentityProvider {
	return &singleIdentityProvider{}
}

func (p *singleIdentityProvider) GetUserIdentityForCluster(username string, cluster *fedv1.Cluster) (*identityprovider.Identity, error) {
	return &identityprovider.Identity{
		// TODO: this is supposed to change that the secrets will be stored in federation instead of hosting cluster later
		Location: identityprovider.IdentityLocationHostCluster,
		CredentialRef: corev1.ObjectReference{
			Kind:      "Secret",
			Namespace: os.Getenv("POD_NAMESPACE"),
			Name:      cluster.Spec.SecretRef.Name,
		},
	}, nil
}

func init() {
	identityprovider.RegisterIdentityProvider(providerName, func(config io.Reader) (identityprovider.Interface, error) {
		return NewSingleIdentityProvider(), nil
	})
}
