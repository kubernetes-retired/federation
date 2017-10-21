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

package validation

import (
	"fmt"
	"net"

	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/federation/apis/federation"
	"k8s.io/kubernetes/pkg/api/validation"
)

func ValidateClusterSpec(spec *federation.ClusterSpec, fieldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	// address is required.
	if len(spec.ServerAddressByClientCIDRs) == 0 {
		allErrs = append(allErrs, field.Required(fieldPath.Child("serverAddressByClientCIDRs"), ""))
	} else {
		for i, address := range spec.ServerAddressByClientCIDRs {
			idxPath := fieldPath.Child("serverAddressByClientCIDRs").Index(i)
			if len(address.ClientCIDR) > 0 {
				if _, _, err := net.ParseCIDR(address.ClientCIDR); err != nil {
					allErrs = append(allErrs, field.Invalid(idxPath.Child("clientCIDR"), address.ClientCIDR, fmt.Sprintf("must be a valid CIDR: %v", err)))
				}
			}
		}
	}
	return allErrs
}

func ValidateCluster(cluster *federation.Cluster) field.ErrorList {
	allErrs := validation.ValidateObjectMeta(&cluster.ObjectMeta, false, validation.ValidateClusterName, field.NewPath("metadata"))
	allErrs = append(allErrs, ValidateClusterSpec(&cluster.Spec, field.NewPath("spec"))...)
	return allErrs
}

func ValidateClusterUpdate(cluster, oldCluster *federation.Cluster) field.ErrorList {
	allErrs := validation.ValidateObjectMetaUpdate(&cluster.ObjectMeta, &oldCluster.ObjectMeta, field.NewPath("metadata"))
	if cluster.Name != oldCluster.Name {
		allErrs = append(allErrs, field.Invalid(field.NewPath("meta", "name"),
			cluster.Name+" != "+oldCluster.Name, "cannot change cluster name"))
	}
	return allErrs
}

func ValidateClusterStatusUpdate(cluster, oldCluster *federation.Cluster) field.ErrorList {
	allErrs := validation.ValidateObjectMetaUpdate(&cluster.ObjectMeta, &oldCluster.ObjectMeta, field.NewPath("metadata"))
	return allErrs
}
