/*
Copyright 2021 The Volcano Authors.

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
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "volcano.sh/apis/pkg/apis/flow/v1alpha1"
	flowv1alpha1 "volcano.sh/apis/pkg/client/applyconfiguration/flow/v1alpha1"
	scheme "volcano.sh/apis/pkg/client/clientset/versioned/scheme"
)

// JobFlowsGetter has a method to return a JobFlowInterface.
// A group's client should implement this interface.
type JobFlowsGetter interface {
	JobFlows(namespace string) JobFlowInterface
}

// JobFlowInterface has methods to work with JobFlow resources.
type JobFlowInterface interface {
	Create(ctx context.Context, jobFlow *v1alpha1.JobFlow, opts v1.CreateOptions) (*v1alpha1.JobFlow, error)
	Update(ctx context.Context, jobFlow *v1alpha1.JobFlow, opts v1.UpdateOptions) (*v1alpha1.JobFlow, error)
	UpdateStatus(ctx context.Context, jobFlow *v1alpha1.JobFlow, opts v1.UpdateOptions) (*v1alpha1.JobFlow, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.JobFlow, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.JobFlowList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.JobFlow, err error)
	Apply(ctx context.Context, jobFlow *flowv1alpha1.JobFlowApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.JobFlow, err error)
	ApplyStatus(ctx context.Context, jobFlow *flowv1alpha1.JobFlowApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.JobFlow, err error)
	JobFlowExpansion
}

// jobFlows implements JobFlowInterface
type jobFlows struct {
	client rest.Interface
	ns     string
}

// newJobFlows returns a JobFlows
func newJobFlows(c *FlowV1alpha1Client, namespace string) *jobFlows {
	return &jobFlows{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the jobFlow, and returns the corresponding jobFlow object, and an error if there is any.
func (c *jobFlows) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.JobFlow, err error) {
	result = &v1alpha1.JobFlow{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("jobflows").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of JobFlows that match those selectors.
func (c *jobFlows) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.JobFlowList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.JobFlowList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("jobflows").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested jobFlows.
func (c *jobFlows) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("jobflows").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a jobFlow and creates it.  Returns the server's representation of the jobFlow, and an error, if there is any.
func (c *jobFlows) Create(ctx context.Context, jobFlow *v1alpha1.JobFlow, opts v1.CreateOptions) (result *v1alpha1.JobFlow, err error) {
	result = &v1alpha1.JobFlow{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("jobflows").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(jobFlow).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a jobFlow and updates it. Returns the server's representation of the jobFlow, and an error, if there is any.
func (c *jobFlows) Update(ctx context.Context, jobFlow *v1alpha1.JobFlow, opts v1.UpdateOptions) (result *v1alpha1.JobFlow, err error) {
	result = &v1alpha1.JobFlow{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("jobflows").
		Name(jobFlow.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(jobFlow).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *jobFlows) UpdateStatus(ctx context.Context, jobFlow *v1alpha1.JobFlow, opts v1.UpdateOptions) (result *v1alpha1.JobFlow, err error) {
	result = &v1alpha1.JobFlow{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("jobflows").
		Name(jobFlow.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(jobFlow).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the jobFlow and deletes it. Returns an error if one occurs.
func (c *jobFlows) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("jobflows").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *jobFlows) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("jobflows").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched jobFlow.
func (c *jobFlows) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.JobFlow, err error) {
	result = &v1alpha1.JobFlow{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("jobflows").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied jobFlow.
func (c *jobFlows) Apply(ctx context.Context, jobFlow *flowv1alpha1.JobFlowApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.JobFlow, err error) {
	if jobFlow == nil {
		return nil, fmt.Errorf("jobFlow provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(jobFlow)
	if err != nil {
		return nil, err
	}
	name := jobFlow.Name
	if name == nil {
		return nil, fmt.Errorf("jobFlow.Name must be provided to Apply")
	}
	result = &v1alpha1.JobFlow{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("jobflows").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *jobFlows) ApplyStatus(ctx context.Context, jobFlow *flowv1alpha1.JobFlowApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.JobFlow, err error) {
	if jobFlow == nil {
		return nil, fmt.Errorf("jobFlow provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(jobFlow)
	if err != nil {
		return nil, err
	}

	name := jobFlow.Name
	if name == nil {
		return nil, fmt.Errorf("jobFlow.Name must be provided to Apply")
	}

	result = &v1alpha1.JobFlow{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("jobflows").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
