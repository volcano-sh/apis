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

package v1beta1

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	v1beta1 "volcano.sh/apis/v2/pkg/apis/scheduling/v1beta1"
	schedulingv1beta1 "volcano.sh/apis/v2/pkg/client/applyconfiguration/scheduling/v1beta1"
	scheme "volcano.sh/apis/v2/pkg/client/clientset/versioned/scheme"
)

// PodGroupsGetter has a method to return a PodGroupInterface.
// A group's client should implement this interface.
type PodGroupsGetter interface {
	PodGroups(namespace string) PodGroupInterface
}

// PodGroupInterface has methods to work with PodGroup resources.
type PodGroupInterface interface {
	Create(ctx context.Context, podGroup *v1beta1.PodGroup, opts v1.CreateOptions) (*v1beta1.PodGroup, error)
	Update(ctx context.Context, podGroup *v1beta1.PodGroup, opts v1.UpdateOptions) (*v1beta1.PodGroup, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, podGroup *v1beta1.PodGroup, opts v1.UpdateOptions) (*v1beta1.PodGroup, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.PodGroup, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.PodGroupList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.PodGroup, err error)
	Apply(ctx context.Context, podGroup *schedulingv1beta1.PodGroupApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.PodGroup, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, podGroup *schedulingv1beta1.PodGroupApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.PodGroup, err error)
	PodGroupExpansion
}

// podGroups implements PodGroupInterface
type podGroups struct {
	*gentype.ClientWithListAndApply[*v1beta1.PodGroup, *v1beta1.PodGroupList, *schedulingv1beta1.PodGroupApplyConfiguration]
}

// newPodGroups returns a PodGroups
func newPodGroups(c *SchedulingV1beta1Client, namespace string) *podGroups {
	return &podGroups{
		gentype.NewClientWithListAndApply[*v1beta1.PodGroup, *v1beta1.PodGroupList, *schedulingv1beta1.PodGroupApplyConfiguration](
			"podgroups",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1beta1.PodGroup { return &v1beta1.PodGroup{} },
			func() *v1beta1.PodGroupList { return &v1beta1.PodGroupList{} }),
	}
}
