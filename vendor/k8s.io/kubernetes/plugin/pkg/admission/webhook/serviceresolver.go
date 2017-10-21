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

// Package webhook checks a webhook for configured operation admission
package webhook

import (
	"errors"
	"fmt"
	"net/url"

	admissioninit "k8s.io/kubernetes/pkg/kubeapiserver/admission"
)

type defaultServiceResolver struct{}

var _ admissioninit.ServiceResolver = defaultServiceResolver{}

// ResolveEndpoint constructs a service URL from a given namespace and name
// note that the name and namespace are required and by default all created addresses use HTTPS scheme.
// for example:
//  name=ross namespace=andromeda resolves to https://ross.andromeda.svc
func (sr defaultServiceResolver) ResolveEndpoint(namespace, name string) (*url.URL, error) {
	if len(name) == 0 || len(namespace) == 0 {
		return &url.URL{}, errors.New("cannot resolve an empty service name or namespace")
	}

	return url.Parse(fmt.Sprintf("https://%s.%s.svc", name, namespace))
}
