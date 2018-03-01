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

package framework

import (
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	clientset "k8s.io/client-go/kubernetes"
)

const (
	// Time required by the loadbalancer to cleanup, proportional to numApps/Ing.
	// Bring the cleanup timeout back down to 5m once b/33588344 is resolved.
	LoadBalancerCleanupTimeout = 15 * time.Minute

	// On average it takes ~6 minutes for a single backend to come online in GCE.
	LoadBalancerPollTimeout  = 15 * time.Minute
	LoadBalancerPollInterval = 30 * time.Second
)

func CleanupServiceResources(c clientset.Interface, loadBalancerName, zone string) {
	if TestContext.Provider == "gce" || TestContext.Provider == "gke" {
		CleanupServiceGCEResources(c, loadBalancerName, zone)
	}

	// TODO: we need to add this function with other cloud providers, if there is a need.
}

func CleanupServiceGCEResources(c clientset.Interface, loadBalancerName, zone string) {
	if pollErr := wait.Poll(5*time.Second, LoadBalancerCleanupTimeout, func() (bool, error) {
		if err := CleanupGCEResources(c, loadBalancerName, zone); err != nil {
			Logf("Still waiting for glbc to cleanup: %v", err)
			return false, nil
		}
		return true, nil
	}); pollErr != nil {
		Failf("Failed to cleanup service GCE resources.")
	}
}
