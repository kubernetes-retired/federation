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

package framework

import (
	"flag"

	"k8s.io/kubernetes/test/e2e/framework"
)

type FederationTestContextType struct {
	// Federation e2e context
	FederatedKubeContext string
	// Federation control plane version to upgrade to while doing upgrade tests
	FederationUpgradeTarget string
	// Whether configuration for accessing federation member clusters should be sourced from the host cluster
	FederationConfigFromCluster bool
}

var TestContext FederationTestContextType

func ViperizeFlags() {
	registerFederationFlags()
	framework.ViperizeFlags()
}

func registerFederationFlags() {
	flag.StringVar(&TestContext.FederatedKubeContext, "federated-kube-context", "e2e-federation", "kubeconfig context for federation.")
	flag.BoolVar(&TestContext.FederationConfigFromCluster, "federation-config-from-cluster", false, "whether to source configuration for member clusters from the hosting cluster.")
	flag.StringVar(&TestContext.FederationUpgradeTarget, "federation-upgrade-target", "ci/latest", "Version to upgrade to (e.g. 'release/stable', 'release/latest', 'ci/latest', '0.19.1', '0.19.1-669-gabac8c8') if doing an federation upgrade test.")
}
