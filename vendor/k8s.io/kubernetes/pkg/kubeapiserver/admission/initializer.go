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

package admission

import (
	"net/http"
	"net/url"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	informers "k8s.io/kubernetes/pkg/client/informers/informers_generated/internalversion"
	"k8s.io/kubernetes/pkg/quota"
)

// TODO add a `WantsToRun` which takes a stopCh.  Might make it generic.

// WantsInternalKubeClientSet defines a function which sets ClientSet for admission plugins that need it
type WantsInternalKubeClientSet interface {
	SetInternalKubeClientSet(internalclientset.Interface)
	admission.Validator
}

// WantsInternalKubeInformerFactory defines a function which sets InformerFactory for admission plugins that need it
type WantsInternalKubeInformerFactory interface {
	SetInternalKubeInformerFactory(informers.SharedInformerFactory)
	admission.Validator
}

// WantsCloudConfig defines a function which sets CloudConfig for admission plugins that need it.
type WantsCloudConfig interface {
	SetCloudConfig([]byte)
}

// WantsRESTMapper defines a function which sets RESTMapper for admission plugins that need it.
type WantsRESTMapper interface {
	SetRESTMapper(meta.RESTMapper)
}

// WantsQuotaRegistry defines a function which sets quota registry for admission plugins that need it.
type WantsQuotaRegistry interface {
	SetQuotaRegistry(quota.Registry)
	admission.Validator
}

// WantsServiceResolver defines a fuction that accepts a ServiceResolver for
// admission plugins that need to make calls to services.
type WantsServiceResolver interface {
	SetServiceResolver(ServiceResolver)
}

// ServiceResolver knows how to convert a service reference into an actual
// location.
type ServiceResolver interface {
	ResolveEndpoint(namespace, name string) (*url.URL, error)
}

// WantsProxyTransport defines a fuction that accepts a proxy transport for admission
// plugins that need to make calls to pods.
type WantsProxyTransport interface {
	SetProxyTransport(proxyTransport *http.Transport)
}

type PluginInitializer struct {
	internalClient  internalclientset.Interface
	externalClient  clientset.Interface
	informers       informers.SharedInformerFactory
	authorizer      authorizer.Authorizer
	cloudConfig     []byte
	restMapper      meta.RESTMapper
	quotaRegistry   quota.Registry
	serviceResolver ServiceResolver

	// for proving we are apiserver in call-outs
	proxyTransport *http.Transport
}

var _ admission.PluginInitializer = &PluginInitializer{}

// NewPluginInitializer constructs new instance of PluginInitializer
// TODO: switch these parameters to use the builder pattern or just make them
// all public, this construction method is pointless boilerplate.
func NewPluginInitializer(
	internalClient internalclientset.Interface,
	sharedInformers informers.SharedInformerFactory,
	cloudConfig []byte,
	restMapper meta.RESTMapper,
	quotaRegistry quota.Registry,
) *PluginInitializer {
	return &PluginInitializer{
		internalClient: internalClient,
		informers:      sharedInformers,
		cloudConfig:    cloudConfig,
		restMapper:     restMapper,
		quotaRegistry:  quotaRegistry,
	}
}

// SetServiceResolver sets the service resolver which is needed by some plugins.
func (i *PluginInitializer) SetServiceResolver(s ServiceResolver) *PluginInitializer {
	i.serviceResolver = s
	return i
}

// SetProxyTransport sets the proxyTransport which is needed by some plugins.
func (i *PluginInitializer) SetProxyTransport(proxyTransport *http.Transport) *PluginInitializer {
	i.proxyTransport = proxyTransport
	return i
}

// Initialize checks the initialization interfaces implemented by each plugin
// and provide the appropriate initialization data
func (i *PluginInitializer) Initialize(plugin admission.Interface) {
	if wants, ok := plugin.(WantsInternalKubeClientSet); ok {
		wants.SetInternalKubeClientSet(i.internalClient)
	}

	if wants, ok := plugin.(WantsInternalKubeInformerFactory); ok {
		wants.SetInternalKubeInformerFactory(i.informers)
	}

	if wants, ok := plugin.(WantsCloudConfig); ok {
		wants.SetCloudConfig(i.cloudConfig)
	}

	if wants, ok := plugin.(WantsRESTMapper); ok {
		wants.SetRESTMapper(i.restMapper)
	}

	if wants, ok := plugin.(WantsQuotaRegistry); ok {
		wants.SetQuotaRegistry(i.quotaRegistry)
	}

	if wants, ok := plugin.(WantsServiceResolver); ok {
		wants.SetServiceResolver(i.serviceResolver)
	}

	if wants, ok := plugin.(WantsProxyTransport); ok {
		wants.SetProxyTransport(i.proxyTransport)
	}
}
