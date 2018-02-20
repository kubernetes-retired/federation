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
	"io"
	"strings"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/authentication/serviceaccount"
)

const (
	// LastUpdatedByUserAnno is the annotation key used to track the last user who updated the object
	LastUpdatedByUserAnno = "federation.alpha.kubernetes.io/last-updated-by-user"
)

// Register registers a plugin
func Register(plugins *admission.Plugins) {
	plugins.Register("LastUpdatedBy", func(config io.Reader) (admission.Interface, error) {
		return NewLastUpdatedByAdmit(), nil
	})
}

// lastUpdatedByAdmit is an implementation of admission.Interface which injects the user who last updated the object.
type lastUpdatedByAdmit struct {
	*admission.Handler
}

func (a *lastUpdatedByAdmit) Admit(attributes admission.Attributes) error {

	// Ignore all calls to subresources
	if len(attributes.GetSubresource()) != 0 {
		return nil
	}

	// Not an API object
	accessor, err := meta.Accessor(attributes.GetObject())
	if err != nil {
		return nil
	}

	// Ignore special objects such as SubjectAccessReview, etc.
	if accessor.GetName() == "" {
		return nil
	}

	username := attributes.GetUserInfo().GetName()
	// Ignore service account calls
	if strings.HasPrefix(username, serviceaccount.ServiceAccountUsernamePrefix) {
		return nil
	}

	annotations := accessor.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
	}
	annotations[LastUpdatedByUserAnno] = username
	accessor.SetAnnotations(annotations)

	return nil
}

// NewLastUpdatedByAdmit creates a new last updated by admit admission handler
func NewLastUpdatedByAdmit() admission.Interface {
	return &lastUpdatedByAdmit{
		Handler: admission.NewHandler(admission.Create, admission.Update),
	}
}
