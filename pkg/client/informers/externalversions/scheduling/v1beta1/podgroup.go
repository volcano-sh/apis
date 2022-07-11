/*
Copyright 2022 The Volcano Authors.

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
// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	schedulingv1beta1 "volcano.sh/apis/pkg/apis/scheduling/v1beta1"
	versioned "volcano.sh/apis/pkg/client/clientset/versioned"
	internalinterfaces "volcano.sh/apis/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "volcano.sh/apis/pkg/client/listers/scheduling/v1beta1"
)

// PodGroupInformer provides access to a shared informer and lister for
// PodGroups.
type PodGroupInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.PodGroupLister
}

type podGroupInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPodGroupInformer constructs a new informer for PodGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPodGroupInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPodGroupInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPodGroupInformer constructs a new informer for PodGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPodGroupInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SchedulingV1beta1().PodGroups(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SchedulingV1beta1().PodGroups(namespace).Watch(context.TODO(), options)
			},
		},
		&schedulingv1beta1.PodGroup{},
		resyncPeriod,
		indexers,
	)
}

func (f *podGroupInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPodGroupInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *podGroupInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&schedulingv1beta1.PodGroup{}, f.defaultInformer)
}

func (f *podGroupInformer) Lister() v1beta1.PodGroupLister {
	return v1beta1.NewPodGroupLister(f.Informer().GetIndexer())
}
