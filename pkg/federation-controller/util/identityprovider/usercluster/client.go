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

package usercluster

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
)

type crdClient struct {
	client         rest.Interface
	parameterCodec runtime.ParameterCodec
}

type userClusterIdentityClient struct {
	client         rest.Interface
	ns             string
	parameterCodec runtime.ParameterCodec
}

// NewCRDClientForConfig creates a client of the CRD object for user-cluster identity provider
func NewCRDClientForConfig(cfg *rest.Config) (*crdClient, *runtime.Scheme, error) {
	crdScheme := runtime.NewScheme()
	if err := AddToScheme(crdScheme); err != nil {
		return nil, nil, err
	}

	config := *cfg
	config.GroupVersion = &SchemeGroupVersion
	config.APIPath = "/apis"
	config.ContentType = runtime.ContentTypeJSON
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: serializer.NewCodecFactory(crdScheme)}

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, nil, err
	}

	return &crdClient{client, runtime.NewParameterCodec(crdScheme)}, crdScheme, nil
}

func (c *crdClient) UserClusterIdentity(namespace string) *userClusterIdentityClient {
	return &userClusterIdentityClient{c.client, namespace, c.parameterCodec}
}

// Get takes name of the event, and returns the corresponding event object, and an error if there is any.
func (c *userClusterIdentityClient) Get(name string, options metav1.GetOptions) (result *UserClusterIdentity, err error) {
	result = &UserClusterIdentity{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource(UserClusterIdentityPlural).
		Name(name).
		VersionedParams(&options, c.parameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Events that match those selectors.
func (c *userClusterIdentityClient) List(opts metav1.ListOptions) (result *UserClusterIdentityList, err error) {
	result = &UserClusterIdentityList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource(UserClusterIdentityPlural).
		VersionedParams(&opts, c.parameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested events.
func (c *userClusterIdentityClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource(UserClusterIdentityPlural).
		VersionedParams(&opts, c.parameterCodec).
		Watch()
}

// Create takes the representation of a event and creates it.  Returns the server's representation of the event, and an error, if there is any.
func (c *userClusterIdentityClient) Create(event *UserClusterIdentity) (result *UserClusterIdentity, err error) {
	result = &UserClusterIdentity{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource(UserClusterIdentityPlural).
		Body(event).
		Do().
		Into(result)
	return
}

// Update takes the representation of a event and updates it. Returns the server's representation of the event, and an error, if there is any.
func (c *userClusterIdentityClient) Update(event *UserClusterIdentity) (result *UserClusterIdentity, err error) {
	result = &UserClusterIdentity{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource(UserClusterIdentityPlural).
		Name(event.Name).
		Body(event).
		Do().
		Into(result)
	return
}

// Delete takes name of the event and deletes it. Returns an error if one occurs.
func (c *userClusterIdentityClient) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource(UserClusterIdentityPlural).
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *userClusterIdentityClient) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource(UserClusterIdentityPlural).
		VersionedParams(&listOptions, c.parameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched event.
func (c *userClusterIdentityClient) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *UserClusterIdentity, err error) {
	result = &UserClusterIdentity{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource(UserClusterIdentityPlural).
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
