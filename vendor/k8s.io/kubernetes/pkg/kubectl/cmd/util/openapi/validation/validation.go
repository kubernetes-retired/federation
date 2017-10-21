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

package validation

import (
	"errors"
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/kubernetes/pkg/kubectl/cmd/util/openapi"
)

type SchemaValidation struct {
	resources openapi.Resources
}

func NewSchemaValidation(resources openapi.Resources) *SchemaValidation {
	return &SchemaValidation{
		resources: resources,
	}
}

func (v *SchemaValidation) ValidateBytes(data []byte) error {
	obj, err := parse(data)
	if err != nil {
		return err
	}

	gvk, err := getObjectKind(obj)
	if err != nil {
		return err
	}

	if strings.HasSuffix(gvk.Kind, "List") {
		return utilerrors.NewAggregate(v.validateList(obj))
	}

	return utilerrors.NewAggregate(v.validateResource(obj, gvk))
}

func (v *SchemaValidation) validateList(object interface{}) []error {
	fields := object.(map[string]interface{})
	if fields == nil {
		return []error{errors.New("invalid object to validate")}
	}

	errs := []error{}
	for _, item := range fields["items"].([]interface{}) {
		if gvk, err := getObjectKind(item); err != nil {
			errs = append(errs, err)
		} else {
			errs = append(errs, v.validateResource(item, gvk)...)
		}
	}
	return errs
}

func (v *SchemaValidation) validateResource(obj interface{}, gvk schema.GroupVersionKind) []error {
	resource := v.resources.LookupResource(gvk)
	if resource == nil {
		// resource is not present, let's just skip validation.
		return nil
	}

	rootValidation, err := itemFactory(openapi.NewPath(gvk.Kind), obj)
	if err != nil {
		return []error{err}
	}
	resource.Accept(rootValidation)
	return rootValidation.Errors()
}

func parse(data []byte) (interface{}, error) {
	var obj interface{}
	out, err := yaml.ToJSON(data)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(out, &obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func getObjectKind(object interface{}) (schema.GroupVersionKind, error) {
	fields := object.(map[string]interface{})
	if fields == nil {
		return schema.GroupVersionKind{}, errors.New("invalid object to validate")
	}
	apiVersion := fields["apiVersion"]
	if apiVersion == nil {
		return schema.GroupVersionKind{}, errors.New("apiVersion not set")
	}
	if _, ok := apiVersion.(string); !ok {
		return schema.GroupVersionKind{}, errors.New("apiVersion isn't string type")
	}
	gv, err := schema.ParseGroupVersion(apiVersion.(string))
	if err != nil {
		return schema.GroupVersionKind{}, err
	}
	kind := fields["kind"]
	if kind == nil {
		return schema.GroupVersionKind{}, errors.New("kind not set")
	}
	if _, ok := kind.(string); !ok {
		return schema.GroupVersionKind{}, errors.New("kind isn't string type")
	}

	return schema.GroupVersionKind{Group: gv.Group, Version: gv.Version, Kind: kind.(string)}, nil
}
