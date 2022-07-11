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
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "volcano.sh/apis/pkg/apis/flow/v1alpha1"
)

// JobFlowLister helps list JobFlows.
// All objects returned here must be treated as read-only.
type JobFlowLister interface {
	// List lists all JobFlows in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.JobFlow, err error)
	// JobFlows returns an object that can list and get JobFlows.
	JobFlows(namespace string) JobFlowNamespaceLister
	JobFlowListerExpansion
}

// jobFlowLister implements the JobFlowLister interface.
type jobFlowLister struct {
	indexer cache.Indexer
}

// NewJobFlowLister returns a new JobFlowLister.
func NewJobFlowLister(indexer cache.Indexer) JobFlowLister {
	return &jobFlowLister{indexer: indexer}
}

// List lists all JobFlows in the indexer.
func (s *jobFlowLister) List(selector labels.Selector) (ret []*v1alpha1.JobFlow, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.JobFlow))
	})
	return ret, err
}

// JobFlows returns an object that can list and get JobFlows.
func (s *jobFlowLister) JobFlows(namespace string) JobFlowNamespaceLister {
	return jobFlowNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// JobFlowNamespaceLister helps list and get JobFlows.
// All objects returned here must be treated as read-only.
type JobFlowNamespaceLister interface {
	// List lists all JobFlows in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.JobFlow, err error)
	// Get retrieves the JobFlow from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.JobFlow, error)
	JobFlowNamespaceListerExpansion
}

// jobFlowNamespaceLister implements the JobFlowNamespaceLister
// interface.
type jobFlowNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all JobFlows in the indexer for a given namespace.
func (s jobFlowNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.JobFlow, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.JobFlow))
	})
	return ret, err
}

// Get retrieves the JobFlow from the indexer for a given namespace and name.
func (s jobFlowNamespaceLister) Get(name string) (*v1alpha1.JobFlow, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("jobflow"), name)
	}
	return obj.(*v1alpha1.JobFlow), nil
}
