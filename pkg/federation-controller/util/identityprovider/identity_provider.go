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

package identityprovider

import (
	authv1 "k8s.io/api/authentication/v1"
	corev1 "k8s.io/api/core/v1"
	fedv1 "k8s.io/federation/apis/federation/v1beta1"
)

// IdentityLocation is where Kubernetes federation control plane to find the referenced credential object
type IdentityLocation string

const (
	// IdentityLocationFederation means the credential object was stored in the federation control plane
	IdentityLocationFederation IdentityLocation = "Federation"
	// IdentityLocationHostCluster means the credential object was stored in the host cluster where federation control plane is running
	IdentityLocationHostCluster IdentityLocation = "HostCluster"
	// IdentityLocationCluster means the credential object was stored in the cluster itself, where federation control plane will need read access to it
	IdentityLocationCluster IdentityLocation = "Cluster"
)

// +k8s:deepcopy-gen=true
// Identity contains the info to connect to a cluster
type Identity struct {
	// Location is where `CredentialRef` object was stored.
	Location IdentityLocation
	// Reference to a service account or a secret (with kubeconfig) containing credentials for federation to use.
	CredentialRef corev1.ObjectReference
	// Information about the user to impersonate, if any.
	ImpersonatingUser *authv1.UserInfo
}

// Interface for a identity provider
type IdentityProvider interface {
	// GetUserIdentityForCluster returns the identity should be used for the user to access the given cluster
	GetUserIdentityForCluster(user string, cluster *fedv1.Cluster) (*Identity, error)
}
