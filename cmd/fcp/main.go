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

// A binary that can morph into all of the other kubernetes binaries. You can
// also soft-link to it busybox style.
//
package main

import (
	"os"

	_ "k8s.io/kubernetes/pkg/client/metrics/prometheus" // for client metric registration
	_ "k8s.io/kubernetes/pkg/version/prometheus"        // for version metric registration
)

func main() {
	fcp := HyperKube{
		Name: "fcp",
		Long: "This is an all-in-one binary that can run any of the various Federation servers.",
	}

	fcp.AddServer(NewFederationAPIServer())
	fcp.AddServer(NewFederationCMServer())

	fcp.RunToExit(os.Args)
}
