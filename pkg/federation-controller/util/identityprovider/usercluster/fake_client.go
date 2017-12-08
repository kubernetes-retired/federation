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
	"k8s.io/client-go/testing"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

var scheme = runtime.NewScheme()
var codecs = serializer.NewCodecFactory(scheme)
var parameterCodec = runtime.NewParameterCodec(scheme)

var resource = SchemeGroupVersion.WithResource(UserClusterIdentityPlural)
var kind = SchemeGroupVersion.WithKind("UserClusterIdentity")

type fakeCRD struct {
	*testing.Fake
}

func (c* fakeCRD) UserClusterIdentity(namespace string) *fakeUserClusterIdentitys {
	return &fakeUserClusterIdentitys{c, namespace}
}

type fakeUserClusterIdentitys struct {
	Fake *fakeCRD
	ns string
}

func init() {
	AddToScheme(scheme)
}

func NewFakeCRDClient(objects ...runtime.Object) *fakeCRD {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	fakePtr := testing.Fake{}
	fakePtr.AddReactor("*", "*", testing.ObjectReaction(o))
	fakePtr.AddWatchReactor("*", testing.DefaultWatchReactor(watch.NewFake(), nil))

	return &fakeCRD{&fakePtr}
}

// Get takes name of the identity, and returns the corresponding UserClusterIdentity object, and an error if there is any.
func (c *fakeUserClusterIdentitys) Get(name string, options metav1.GetOptions) (result *UserClusterIdentity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(resource, c.ns, name), &UserClusterIdentity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*UserClusterIdentity), err
}

// List takes label and field selectors, and returns the list of UserClusterIdentitys that match those selectors.
func (c *fakeUserClusterIdentitys) List(opts metav1.ListOptions) (result *UserClusterIdentityList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(resource, kind, c.ns, opts), &UserClusterIdentityList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &UserClusterIdentityList{}
	for _, item := range obj.(*UserClusterIdentityList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested UserClusterIdentitys.
func (c *fakeUserClusterIdentitys) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(resource, c.ns, opts))

}

// Create takes the representation of a UserClusterIdentity and creates it.  Returns the server's representation of the identity, and an error, if there is any.
func (c *fakeUserClusterIdentitys) Create(identity *UserClusterIdentity) (result *UserClusterIdentity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(resource, c.ns, identity), &UserClusterIdentity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*UserClusterIdentity), err
}

// Update takes the representation of a UserClusterIdentity and updates it. Returns the server's representation of the identity, and an error, if there is any.
func (c *fakeUserClusterIdentitys) Update(identity *UserClusterIdentity) (result *UserClusterIdentity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(resource, c.ns, identity), &UserClusterIdentity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*UserClusterIdentity), err
}

// Delete takes name of the UserClusterIdentity and deletes it. Returns an error if one occurs.
func (c *fakeUserClusterIdentitys) Delete(name string, options *metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(resource, c.ns, name), &UserClusterIdentity{})

	return err
}
