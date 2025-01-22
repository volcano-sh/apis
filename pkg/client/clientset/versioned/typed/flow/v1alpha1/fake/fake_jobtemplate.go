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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "volcano.sh/apis/v2/pkg/apis/flow/v1alpha1"
	flowv1alpha1 "volcano.sh/apis/v2/pkg/client/applyconfiguration/flow/v1alpha1"
)

// FakeJobTemplates implements JobTemplateInterface
type FakeJobTemplates struct {
	Fake *FakeFlowV1alpha1
	ns   string
}

var jobtemplatesResource = v1alpha1.SchemeGroupVersion.WithResource("jobtemplates")

var jobtemplatesKind = v1alpha1.SchemeGroupVersion.WithKind("JobTemplate")

// Get takes name of the jobTemplate, and returns the corresponding jobTemplate object, and an error if there is any.
func (c *FakeJobTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.JobTemplate, err error) {
	emptyResult := &v1alpha1.JobTemplate{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(jobtemplatesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.JobTemplate), err
}

// List takes label and field selectors, and returns the list of JobTemplates that match those selectors.
func (c *FakeJobTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.JobTemplateList, err error) {
	emptyResult := &v1alpha1.JobTemplateList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(jobtemplatesResource, jobtemplatesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.JobTemplateList{ListMeta: obj.(*v1alpha1.JobTemplateList).ListMeta}
	for _, item := range obj.(*v1alpha1.JobTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested jobTemplates.
func (c *FakeJobTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(jobtemplatesResource, c.ns, opts))

}

// Create takes the representation of a jobTemplate and creates it.  Returns the server's representation of the jobTemplate, and an error, if there is any.
func (c *FakeJobTemplates) Create(ctx context.Context, jobTemplate *v1alpha1.JobTemplate, opts v1.CreateOptions) (result *v1alpha1.JobTemplate, err error) {
	emptyResult := &v1alpha1.JobTemplate{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(jobtemplatesResource, c.ns, jobTemplate, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.JobTemplate), err
}

// Update takes the representation of a jobTemplate and updates it. Returns the server's representation of the jobTemplate, and an error, if there is any.
func (c *FakeJobTemplates) Update(ctx context.Context, jobTemplate *v1alpha1.JobTemplate, opts v1.UpdateOptions) (result *v1alpha1.JobTemplate, err error) {
	emptyResult := &v1alpha1.JobTemplate{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(jobtemplatesResource, c.ns, jobTemplate, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.JobTemplate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeJobTemplates) UpdateStatus(ctx context.Context, jobTemplate *v1alpha1.JobTemplate, opts v1.UpdateOptions) (result *v1alpha1.JobTemplate, err error) {
	emptyResult := &v1alpha1.JobTemplate{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(jobtemplatesResource, "status", c.ns, jobTemplate, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.JobTemplate), err
}

// Delete takes name of the jobTemplate and deletes it. Returns an error if one occurs.
func (c *FakeJobTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(jobtemplatesResource, c.ns, name, opts), &v1alpha1.JobTemplate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeJobTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(jobtemplatesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.JobTemplateList{})
	return err
}

// Patch applies the patch and returns the patched jobTemplate.
func (c *FakeJobTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.JobTemplate, err error) {
	emptyResult := &v1alpha1.JobTemplate{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(jobtemplatesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.JobTemplate), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied jobTemplate.
func (c *FakeJobTemplates) Apply(ctx context.Context, jobTemplate *flowv1alpha1.JobTemplateApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.JobTemplate, err error) {
	if jobTemplate == nil {
		return nil, fmt.Errorf("jobTemplate provided to Apply must not be nil")
	}
	data, err := json.Marshal(jobTemplate)
	if err != nil {
		return nil, err
	}
	name := jobTemplate.Name
	if name == nil {
		return nil, fmt.Errorf("jobTemplate.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.JobTemplate{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(jobtemplatesResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.JobTemplate), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeJobTemplates) ApplyStatus(ctx context.Context, jobTemplate *flowv1alpha1.JobTemplateApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.JobTemplate, err error) {
	if jobTemplate == nil {
		return nil, fmt.Errorf("jobTemplate provided to Apply must not be nil")
	}
	data, err := json.Marshal(jobTemplate)
	if err != nil {
		return nil, err
	}
	name := jobTemplate.Name
	if name == nil {
		return nil, fmt.Errorf("jobTemplate.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.JobTemplate{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(jobtemplatesResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.JobTemplate), err
}
