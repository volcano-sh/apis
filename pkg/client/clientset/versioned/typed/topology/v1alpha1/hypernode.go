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

package v1alpha1

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	v1alpha1 "volcano.sh/apis/pkg/apis/topology/v1alpha1"
	topologyv1alpha1 "volcano.sh/apis/pkg/client/applyconfiguration/topology/v1alpha1"
	scheme "volcano.sh/apis/pkg/client/clientset/versioned/scheme"
)

// HyperNodesGetter has a method to return a HyperNodeInterface.
// A group's client should implement this interface.
type HyperNodesGetter interface {
	HyperNodes() HyperNodeInterface
}

// HyperNodeInterface has methods to work with HyperNode resources.
type HyperNodeInterface interface {
	Create(ctx context.Context, hyperNode *v1alpha1.HyperNode, opts v1.CreateOptions) (*v1alpha1.HyperNode, error)
	Update(ctx context.Context, hyperNode *v1alpha1.HyperNode, opts v1.UpdateOptions) (*v1alpha1.HyperNode, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, hyperNode *v1alpha1.HyperNode, opts v1.UpdateOptions) (*v1alpha1.HyperNode, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.HyperNode, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.HyperNodeList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HyperNode, err error)
	Apply(ctx context.Context, hyperNode *topologyv1alpha1.HyperNodeApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.HyperNode, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, hyperNode *topologyv1alpha1.HyperNodeApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.HyperNode, err error)
	HyperNodeExpansion
}

// hyperNodes implements HyperNodeInterface
type hyperNodes struct {
	*gentype.ClientWithListAndApply[*v1alpha1.HyperNode, *v1alpha1.HyperNodeList, *topologyv1alpha1.HyperNodeApplyConfiguration]
}

// newHyperNodes returns a HyperNodes
func newHyperNodes(c *TopologyV1alpha1Client) *hyperNodes {
	return &hyperNodes{
		gentype.NewClientWithListAndApply[*v1alpha1.HyperNode, *v1alpha1.HyperNodeList, *topologyv1alpha1.HyperNodeApplyConfiguration](
			"hypernodes",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1alpha1.HyperNode { return &v1alpha1.HyperNode{} },
			func() *v1alpha1.HyperNodeList { return &v1alpha1.HyperNodeList{} }),
	}
}
