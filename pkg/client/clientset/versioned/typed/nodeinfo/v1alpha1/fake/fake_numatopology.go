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
	v1alpha1 "volcano.sh/apis/pkg/apis/nodeinfo/v1alpha1"
	nodeinfov1alpha1 "volcano.sh/apis/pkg/client/applyconfiguration/nodeinfo/v1alpha1"
)

// FakeNumatopologies implements NumatopologyInterface
type FakeNumatopologies struct {
	Fake *FakeNodeinfoV1alpha1
}

var numatopologiesResource = v1alpha1.SchemeGroupVersion.WithResource("numatopologies")

var numatopologiesKind = v1alpha1.SchemeGroupVersion.WithKind("Numatopology")

// Get takes name of the numatopology, and returns the corresponding numatopology object, and an error if there is any.
func (c *FakeNumatopologies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Numatopology, err error) {
	emptyResult := &v1alpha1.Numatopology{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(numatopologiesResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Numatopology), err
}

// List takes label and field selectors, and returns the list of Numatopologies that match those selectors.
func (c *FakeNumatopologies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NumatopologyList, err error) {
	emptyResult := &v1alpha1.NumatopologyList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(numatopologiesResource, numatopologiesKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NumatopologyList{ListMeta: obj.(*v1alpha1.NumatopologyList).ListMeta}
	for _, item := range obj.(*v1alpha1.NumatopologyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested numatopologies.
func (c *FakeNumatopologies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(numatopologiesResource, opts))
}

// Create takes the representation of a numatopology and creates it.  Returns the server's representation of the numatopology, and an error, if there is any.
func (c *FakeNumatopologies) Create(ctx context.Context, numatopology *v1alpha1.Numatopology, opts v1.CreateOptions) (result *v1alpha1.Numatopology, err error) {
	emptyResult := &v1alpha1.Numatopology{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(numatopologiesResource, numatopology, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Numatopology), err
}

// Update takes the representation of a numatopology and updates it. Returns the server's representation of the numatopology, and an error, if there is any.
func (c *FakeNumatopologies) Update(ctx context.Context, numatopology *v1alpha1.Numatopology, opts v1.UpdateOptions) (result *v1alpha1.Numatopology, err error) {
	emptyResult := &v1alpha1.Numatopology{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(numatopologiesResource, numatopology, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Numatopology), err
}

// Delete takes name of the numatopology and deletes it. Returns an error if one occurs.
func (c *FakeNumatopologies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(numatopologiesResource, name, opts), &v1alpha1.Numatopology{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNumatopologies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(numatopologiesResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NumatopologyList{})
	return err
}

// Patch applies the patch and returns the patched numatopology.
func (c *FakeNumatopologies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Numatopology, err error) {
	emptyResult := &v1alpha1.Numatopology{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(numatopologiesResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Numatopology), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied numatopology.
func (c *FakeNumatopologies) Apply(ctx context.Context, numatopology *nodeinfov1alpha1.NumatopologyApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Numatopology, err error) {
	if numatopology == nil {
		return nil, fmt.Errorf("numatopology provided to Apply must not be nil")
	}
	data, err := json.Marshal(numatopology)
	if err != nil {
		return nil, err
	}
	name := numatopology.Name
	if name == nil {
		return nil, fmt.Errorf("numatopology.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.Numatopology{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(numatopologiesResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Numatopology), err
}
