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
	context "context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	nodeinfov1alpha1 "volcano.sh/apis/pkg/apis/nodeinfo/v1alpha1"
	applyconfigurationnodeinfov1alpha1 "volcano.sh/apis/pkg/client/applyconfiguration/nodeinfo/v1alpha1"
	scheme "volcano.sh/apis/pkg/client/clientset/versioned/scheme"
)

// NumatopologiesGetter has a method to return a NumatopologyInterface.
// A group's client should implement this interface.
type NumatopologiesGetter interface {
	Numatopologies() NumatopologyInterface
}

// NumatopologyInterface has methods to work with Numatopology resources.
type NumatopologyInterface interface {
	Create(ctx context.Context, numatopology *nodeinfov1alpha1.Numatopology, opts v1.CreateOptions) (*nodeinfov1alpha1.Numatopology, error)
	Update(ctx context.Context, numatopology *nodeinfov1alpha1.Numatopology, opts v1.UpdateOptions) (*nodeinfov1alpha1.Numatopology, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*nodeinfov1alpha1.Numatopology, error)
	List(ctx context.Context, opts v1.ListOptions) (*nodeinfov1alpha1.NumatopologyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *nodeinfov1alpha1.Numatopology, err error)
	Apply(ctx context.Context, numatopology *applyconfigurationnodeinfov1alpha1.NumatopologyApplyConfiguration, opts v1.ApplyOptions) (result *nodeinfov1alpha1.Numatopology, err error)
	NumatopologyExpansion
}

// numatopologies implements NumatopologyInterface
type numatopologies struct {
	*gentype.ClientWithListAndApply[*nodeinfov1alpha1.Numatopology, *nodeinfov1alpha1.NumatopologyList, *applyconfigurationnodeinfov1alpha1.NumatopologyApplyConfiguration]
}

// newNumatopologies returns a Numatopologies
func newNumatopologies(c *NodeinfoV1alpha1Client) *numatopologies {
	return &numatopologies{
		gentype.NewClientWithListAndApply[*nodeinfov1alpha1.Numatopology, *nodeinfov1alpha1.NumatopologyList, *applyconfigurationnodeinfov1alpha1.NumatopologyApplyConfiguration](
			"numatopologies",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *nodeinfov1alpha1.Numatopology { return &nodeinfov1alpha1.Numatopology{} },
			func() *nodeinfov1alpha1.NumatopologyList { return &nodeinfov1alpha1.NumatopologyList{} },
		),
	}
}
