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

package testapi

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/federation/apis/federation"
	"k8s.io/kubernetes/pkg/api"
	kubetestapi "k8s.io/kubernetes/pkg/api/testapi"

	_ "k8s.io/federation/apis/federation/install"
)

var Federation kubetestapi.TestGroup

func init() {
	if _, ok := kubetestapi.Groups[federation.GroupName]; !ok {
		externalGroupVersion := schema.GroupVersion{Group: federation.GroupName, Version: api.Registry.GroupOrDie(federation.GroupName).GroupVersion.Version}
		kubetestapi.Groups[federation.GroupName] = kubetestapi.NewTestGroup(
			externalGroupVersion,
			federation.SchemeGroupVersion,
			api.Scheme.KnownTypes(federation.SchemeGroupVersion),
			api.Scheme.KnownTypes(externalGroupVersion),
		)
	}
	Federation = kubetestapi.Groups[federation.GroupName]
}
