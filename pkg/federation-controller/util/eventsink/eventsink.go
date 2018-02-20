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

package eventsink

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	fedclientset "k8s.io/federation/client/clientset_generated/federation_clientset"
)

// Implements k8s.io/client-go/tools/record.EventSink.
type FederatedEventSink struct {
	clientset fedclientset.Interface
}

// To check if all required functions are implemented.
var _ record.EventSink = &FederatedEventSink{}

func NewFederatedEventSink(clientset fedclientset.Interface) *FederatedEventSink {
	return &FederatedEventSink{
		clientset: clientset,
	}
}

func (fes *FederatedEventSink) Create(event *v1.Event) (*v1.Event, error) {
	return fes.clientset.Core().Events(event.Namespace).Create(event)
}

func (fes *FederatedEventSink) Update(event *v1.Event) (*v1.Event, error) {
	return fes.clientset.Core().Events(event.Namespace).Update(event)
}

func (fes *FederatedEventSink) Patch(event *v1.Event, data []byte) (*v1.Event, error) {
	return fes.clientset.Core().Events(event.Namespace).Patch(event.Name, types.StrategicMergePatchType, data)
}
