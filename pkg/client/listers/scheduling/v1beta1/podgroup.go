/*
Copyright The Volcano Authors.

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
// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
	v1beta1 "volcano.sh/apis/pkg/apis/scheduling/v1beta1"
)

// PodGroupLister helps list PodGroups.
// All objects returned here must be treated as read-only.
type PodGroupLister interface {
	// List lists all PodGroups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.PodGroup, err error)
	// PodGroups returns an object that can list and get PodGroups.
	PodGroups(namespace string) PodGroupNamespaceLister
	PodGroupListerExpansion
}

// podGroupLister implements the PodGroupLister interface.
type podGroupLister struct {
	listers.ResourceIndexer[*v1beta1.PodGroup]
}

// NewPodGroupLister returns a new PodGroupLister.
func NewPodGroupLister(indexer cache.Indexer) PodGroupLister {
	return &podGroupLister{listers.New[*v1beta1.PodGroup](indexer, v1beta1.Resource("podgroup"))}
}

// PodGroups returns an object that can list and get PodGroups.
func (s *podGroupLister) PodGroups(namespace string) PodGroupNamespaceLister {
	return podGroupNamespaceLister{listers.NewNamespaced[*v1beta1.PodGroup](s.ResourceIndexer, namespace)}
}

// PodGroupNamespaceLister helps list and get PodGroups.
// All objects returned here must be treated as read-only.
type PodGroupNamespaceLister interface {
	// List lists all PodGroups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.PodGroup, err error)
	// Get retrieves the PodGroup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.PodGroup, error)
	PodGroupNamespaceListerExpansion
}

// podGroupNamespaceLister implements the PodGroupNamespaceLister
// interface.
type podGroupNamespaceLister struct {
	listers.ResourceIndexer[*v1beta1.PodGroup]
}
