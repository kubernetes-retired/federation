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

package common

import (
	"fmt"
	"time"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apiserver/pkg/storage/names"
	federationclientset "k8s.io/federation/client/clientset_generated/federation_clientset"
)

const (
	namespacePrefix      = "namespace-test-"
	replicaSetNamePrefix = "namespace-test-rs-"
	DefaultWaitInterval  = 50 * time.Millisecond
)

func CheckNamespaceContentsRemoved(client federationclientset.Interface, l TestLogger) {
	//loggers must be initialized in their respective tests

	// Create namespace
	ns := apiv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: names.SimpleNameGenerator.GenerateName(namespacePrefix),
		},
	}
	l.Logf(fmt.Sprintf("Creating namespace %s", ns.Name))
	_, err := client.Core().Namespaces().Create(&ns)
	if err != nil {
		l.Fatalf("Failed to create namespace %s", ns.Name)
	}
	l.Logf(fmt.Sprintf("Created namespace %s", ns.Name))

	rsName := names.SimpleNameGenerator.GenerateName(replicaSetNamePrefix)
	replicaCount := int32(2)
	rs := &v1beta1.ReplicaSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      rsName,
			Namespace: ns.Name,
		},
		Spec: v1beta1.ReplicaSetSpec{
			Replicas: &replicaCount,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"name": "myrs"},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"name": "myrs"},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "nginx",
							Image: "nginx",
						},
					},
				},
			},
		},
	}
	l.Logf(fmt.Sprintf("Creating replicaset %s in namespace %s", rsName, ns.Name))
	_, err = client.Extensions().ReplicaSets(ns.Name).Create(rs)
	if err != nil {
		l.Fatalf("Failed to create replicaset %v in namespace %s, err: %s", rs, ns.Name, err)
	}
	l.Logf(fmt.Sprintf("Deleting namespace %s", ns.Name))
	deleter := client.Core().Namespaces().Delete
	orphanDeletion := metav1.DeletePropagationOrphan
	err = deleter(ns.Name, &metav1.DeleteOptions{PropagationPolicy: &orphanDeletion})
	if err != nil {
		l.Fatalf("Failed to set %s for deletion: %v", ns.Name, err)
	}
	getter := client.Core().Namespaces().Get
	waitForNamespaceDeletion(ns.Name, l, getter)
}

func waitForNamespaceDeletion(namespace string, l TestLogger, getter func(name string, options metav1.GetOptions) (*apiv1.Namespace, error)) {
	err := wait.PollImmediate(DefaultWaitInterval, wait.ForeverTestTimeout, func() (bool, error) {
		_, err := getter(namespace, metav1.GetOptions{})
		if errors.IsNotFound(err) {
			return true, nil
		} else if err != nil {
			return false, err
		}
		return false, nil
	})
	if err != nil {
		l.Fatalf("Namespaces not deleted: %v", err)
	}
}
