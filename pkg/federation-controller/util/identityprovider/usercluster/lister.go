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
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

const (
	resyncPeriod       = 10 * time.Minute
	userAtClusterIndex = "userAtCluster"
)

type userClusterIdentityInterface interface {
	List(opts metav1.ListOptions) (result *UserClusterIdentityList, err error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
}

type userClusterIdentityLister struct {
	indexer   cache.Indexer
	namespace string
}

func NewUserClusterIdentityLister(client userClusterIdentityInterface, namespace string) *userClusterIdentityLister {
	indexers := cache.Indexers{
		userAtClusterIndex: func(obj interface{}) ([]string, error) {
			identity, ok := obj.(*UserClusterIdentity)
			if !ok {
				return nil, fmt.Errorf("not a UserClusterIdentity object")
			}
			return []string{identity.Spec.Username + "@" + identity.Spec.ClusterName}, nil
		},
	}
	informer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(opts metav1.ListOptions) (runtime.Object, error) {
				return client.List(opts)
			},
			WatchFunc: func(opts metav1.ListOptions) (watch.Interface, error) {
				return client.Watch(opts)
			},
		},
		&UserClusterIdentity{},
		resyncPeriod,
		indexers,
	)

	go informer.Run(wait.NeverStop)
	cache.WaitForCacheSync(wait.NeverStop, informer.HasSynced)

	return &userClusterIdentityLister{
		indexer:   informer.GetIndexer(),
		namespace: namespace,
	}
}

func (l *userClusterIdentityLister) Get(name string) (*UserClusterIdentity, error) {
	obj, exists, err := l.indexer.GetByKey(l.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(Resource("rolebinding"), name)
	}
	return obj.(*UserClusterIdentity), nil
}

func (l *userClusterIdentityLister) ListByUserCluster(username, clusterName string) (ret []*UserClusterIdentity, err error) {
	objs, err := l.indexer.ByIndex(userAtClusterIndex, username+"@"+clusterName)
	if err != nil {
		return nil, err
	}
	for i := range objs {
		ret = append(ret, objs[i].(*UserClusterIdentity))
	}
	return ret, nil
}
