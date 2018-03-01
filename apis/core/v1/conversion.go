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

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/apis/core/v1"
)

func addConversionFuncs(scheme *runtime.Scheme) error {
	// Add non-generated conversion functions
	err := scheme.AddConversionFuncs(
		v1.Convert_v1_DeleteOptions_To_core_DeleteOptions,
		v1.Convert_core_DeleteOptions_To_v1_DeleteOptions,
		v1.Convert_v1_List_To_core_List,
		v1.Convert_core_List_To_v1_List,
		v1.Convert_v1_ListOptions_To_core_ListOptions,
		v1.Convert_core_ListOptions_To_v1_ListOptions,
		v1.Convert_v1_ObjectFieldSelector_To_core_ObjectFieldSelector,
		v1.Convert_core_ObjectFieldSelector_To_v1_ObjectFieldSelector,
		v1.Convert_v1_ObjectMeta_To_core_ObjectMeta,
		v1.Convert_core_ObjectMeta_To_v1_ObjectMeta,
		v1.Convert_v1_ObjectReference_To_core_ObjectReference,
		v1.Convert_core_ObjectReference_To_v1_ObjectReference,
		v1.Convert_v1_Secret_To_core_Secret,
		v1.Convert_core_Secret_To_v1_Secret,
		v1.Convert_v1_SecretList_To_core_SecretList,
		v1.Convert_core_SecretList_To_v1_SecretList,
		v1.Convert_v1_Service_To_core_Service,
		v1.Convert_core_Service_To_v1_Service,
		v1.Convert_v1_ServiceList_To_core_ServiceList,
		v1.Convert_core_ServiceList_To_v1_ServiceList,
		v1.Convert_v1_ServicePort_To_core_ServicePort,
		v1.Convert_core_ServicePort_To_v1_ServicePort,
		v1.Convert_v1_ServiceProxyOptions_To_core_ServiceProxyOptions,
		v1.Convert_core_ServiceProxyOptions_To_v1_ServiceProxyOptions,
		v1.Convert_v1_ServiceSpec_To_core_ServiceSpec,
		v1.Convert_core_ServiceSpec_To_v1_ServiceSpec,
		v1.Convert_v1_ServiceStatus_To_core_ServiceStatus,
		v1.Convert_core_ServiceStatus_To_v1_ServiceStatus,
	)
	if err != nil {
		return err
	}

	if err := v1.AddFieldLabelConversionsForEvent(scheme); err != nil {
		return nil
	}
	if err := v1.AddFieldLabelConversionsForNamespace(scheme); err != nil {
		return nil
	}
	if err := v1.AddFieldLabelConversionsForSecret(scheme); err != nil {
		return nil
	}
	return nil
}
