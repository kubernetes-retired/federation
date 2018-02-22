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

package lastupdatedby

import (
	"testing"

	"github.com/stretchr/testify/assert"

	appsv1b2 "k8s.io/api/apps/v1beta2"
	authnv1 "k8s.io/api/authorization/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/authentication/user"
)

func TestAdmission(t *testing.T) {

	tests := []struct {
		operation   admission.Operation
		handles     bool
		obj         runtime.Object
		oldObj      runtime.Object
		resource    schema.GroupVersionResource
		subResource string
		userInfo    user.Info
		expectUser  string
	}{
		{
			admission.Delete,
			false,
			nil,
			nil,
			schema.GroupVersionResource{},
			"",
			nil,
			"",
		},
		{
			admission.Connect,
			false,
			nil,
			nil,
			schema.GroupVersionResource{},
			"",
			nil,
			"",
		},
		{
			admission.Create,
			true,
			&appsv1b2.Deployment{ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: "ns"}},
			nil,
			appsv1b2.SchemeGroupVersion.WithResource("deployments"),
			"",
			&user.DefaultInfo{Name: "user1"},
			"user1",
		},
		{
			admission.Update,
			true,
			&appsv1b2.Deployment{ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: "ns"}},
			&appsv1b2.Deployment{ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: "ns"}},
			appsv1b2.SchemeGroupVersion.WithResource("deployments"),
			"",
			&user.DefaultInfo{Name: "user1"},
			"user1",
		},
		{
			admission.Create,
			true,
			&appsv1b2.Deployment{ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: "ns"}},
			nil,
			appsv1b2.SchemeGroupVersion.WithResource("deployments"),
			"status",
			&user.DefaultInfo{Name: "user1"},
			"",
		},
		{
			admission.Update,
			true,
			&appsv1b2.Deployment{ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: "ns"}},
			&appsv1b2.Deployment{ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: "ns"}},
			appsv1b2.SchemeGroupVersion.WithResource("deployments"),
			"status",
			&user.DefaultInfo{Name: "user1"},
			"",
		},
		{
			admission.Create,
			true,
			&appsv1b2.Deployment{ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: "ns"}},
			nil,
			appsv1b2.SchemeGroupVersion.WithResource("deployments"),
			"",
			&user.DefaultInfo{Name: "system:serviceaccount:xxx-controller"},
			"",
		},
		{
			admission.Update,
			true,
			&appsv1b2.Deployment{ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: "ns"}},
			&appsv1b2.Deployment{ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: "ns"}},
			appsv1b2.SchemeGroupVersion.WithResource("deployments"),
			"",
			&user.DefaultInfo{Name: "system:serviceaccount:xxx-controller"},
			"",
		},
		{
			admission.Create,
			true,
			&authnv1.SubjectAccessReview{},
			nil,
			authnv1.SchemeGroupVersion.WithResource("subjectaccessreviews"),
			"",
			&user.DefaultInfo{Name: "user1"},
			"",
		},
		{
			admission.Update,
			true,
			&authnv1.SubjectAccessReview{},
			&authnv1.SubjectAccessReview{},
			authnv1.SchemeGroupVersion.WithResource("subjectaccessreviews"),
			"",
			&user.DefaultInfo{Name: "user1"},
			"",
		},
	}

	for _, test := range tests {
		admi := NewLastUpdatedByAdmit()

		handles := admi.Handles(test.operation)
		assert.Equal(t, test.handles, handles)

		if !handles {
			continue
		}

		accessor, _ := meta.Accessor(test.obj)
		attrs := admission.NewAttributesRecord(
			test.obj,
			test.oldObj,
			test.obj.GetObjectKind().GroupVersionKind(),
			accessor.GetNamespace(),
			accessor.GetName(),
			test.resource,
			test.subResource,
			test.operation,
			test.userInfo)

		err := admi.Admit(attrs)
		assert.Nil(t, err)
		assert.Equal(t, test.expectUser, accessor.GetAnnotations()[LastUpdatedByUserAnno])
	}
}
