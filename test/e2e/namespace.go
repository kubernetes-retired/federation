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

package e2e

import (
	"fmt"

	"k8s.io/federation/test/common"
	fedframework "k8s.io/federation/test/e2e/framework"
	"k8s.io/federation/test/k8s/e2e/framework"

	. "github.com/onsi/ginkgo"
)

var _ = framework.KubeDescribe("Federated namespace [Feature:Federation][Experimental]", func() {
	f := fedframework.NewDefaultFederatedFramework("federated-namespace")
	Describe(fmt.Sprintf("Federated namespace resources"), func() {
		It("should delete replicasets in the namespace when the namespace is deleted", func() {
			clientset := f.FederationClientset
			logger := &fedframework.E2eTestLogger{}
			common.CheckNamespaceContentsRemoved(clientset, logger)
		})
	})
})
