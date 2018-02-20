/*
Copyright 2014 The Kubernetes Authors.

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

// This file exists to force the desired plugin implementations to be linked.
// This should probably be part of some configuration fed into the build for a
// given binary target.
import (
	// Cloud providers
	_ "k8s.io/kubernetes/pkg/cloudprovider/providers"

	// Admission policies
	"k8s.io/apiserver/pkg/admission/plugin/initialization"
	"k8s.io/apiserver/pkg/admission/plugin/namespace/lifecycle"
	mutatingwebhook "k8s.io/apiserver/pkg/admission/plugin/webhook/mutating"
	validatingwebhook "k8s.io/apiserver/pkg/admission/plugin/webhook/validating"
	genericoptions "k8s.io/apiserver/pkg/server/options"
	"k8s.io/federation/plugin/pkg/admission/schedulingpolicy"
	"k8s.io/kubernetes/plugin/pkg/admission/gc"
)

// AllOrderedPlugins is the list of all the plugins in order.
var AllOrderedPlugins = []string{
	lifecycle.PluginName,         // NamespaceLifecycle
	gc.PluginName,                // OwnerReferencesPermissionEnforcement
	mutatingwebhook.PluginName,   // MutatingAdmissionWebhook
	initialization.PluginName,    // Initializers
	validatingwebhook.PluginName, // ValidatingAdmissionWebhook
	schedulingpolicy.PluginName,  // SchedulingPolicy
}

// RegisterAllAdmissionPlugins registers all admission plugins
func RegisterAllAdmissionPlugins(admission *genericoptions.AdmissionOptions) {
	gc.Register(admission.Plugins)
	schedulingpolicy.Register(admission.Plugins)

	admission.RecommendedPluginOrder = AllOrderedPlugins
	admission.DefaultOffPlugins.Insert(lifecycle.PluginName, schedulingpolicy.PluginName)
}
