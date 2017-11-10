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

package app

import (
	"github.com/golang/glog"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	"k8s.io/apiserver/pkg/registry/generic"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/apiserver/pkg/server/storage"
	rbacrest "k8s.io/kubernetes/pkg/registry/rbac/rest"

	_ "k8s.io/kubernetes/pkg/apis/rbac/install"
)

func installRBACAPIs(g *genericapiserver.GenericAPIServer, authzer authorizer.Authorizer, optsGetter generic.RESTOptionsGetter, apiResourceConfigSource storage.APIResourceConfigSource) {
	rbacStorageProvider := &rbacrest.RESTStorageProvider{Authorizer: authzer}

	apiGroupInfo, enabled := rbacStorageProvider.NewRESTStorage(apiResourceConfigSource, optsGetter)
	if !enabled {
		glog.Infof("RBAC API not enabled")
	}

	name, hook, err := rbacStorageProvider.PostStartHook()
	if err != nil {
		glog.Fatalf("Error building RBAC PostStartHook: %v", err)
	}
	g.AddPostStartHookOrDie(name, hook)

	if err := g.InstallAPIGroup(&apiGroupInfo); err != nil {
		glog.Fatalf("Error in registering group versions: %v", err)
	}
}
