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
	flowv1alpha1 "volcano.sh/apis/pkg/apis/flow/v1alpha1"
	applyconfigurationflowv1alpha1 "volcano.sh/apis/pkg/client/applyconfiguration/flow/v1alpha1"
	scheme "volcano.sh/apis/pkg/client/clientset/versioned/scheme"
)

// JobTemplatesGetter has a method to return a JobTemplateInterface.
// A group's client should implement this interface.
type JobTemplatesGetter interface {
	JobTemplates(namespace string) JobTemplateInterface
}

// JobTemplateInterface has methods to work with JobTemplate resources.
type JobTemplateInterface interface {
	Create(ctx context.Context, jobTemplate *flowv1alpha1.JobTemplate, opts v1.CreateOptions) (*flowv1alpha1.JobTemplate, error)
	Update(ctx context.Context, jobTemplate *flowv1alpha1.JobTemplate, opts v1.UpdateOptions) (*flowv1alpha1.JobTemplate, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, jobTemplate *flowv1alpha1.JobTemplate, opts v1.UpdateOptions) (*flowv1alpha1.JobTemplate, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*flowv1alpha1.JobTemplate, error)
	List(ctx context.Context, opts v1.ListOptions) (*flowv1alpha1.JobTemplateList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *flowv1alpha1.JobTemplate, err error)
	Apply(ctx context.Context, jobTemplate *applyconfigurationflowv1alpha1.JobTemplateApplyConfiguration, opts v1.ApplyOptions) (result *flowv1alpha1.JobTemplate, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, jobTemplate *applyconfigurationflowv1alpha1.JobTemplateApplyConfiguration, opts v1.ApplyOptions) (result *flowv1alpha1.JobTemplate, err error)
	JobTemplateExpansion
}

// jobTemplates implements JobTemplateInterface
type jobTemplates struct {
	*gentype.ClientWithListAndApply[*flowv1alpha1.JobTemplate, *flowv1alpha1.JobTemplateList, *applyconfigurationflowv1alpha1.JobTemplateApplyConfiguration]
}

// newJobTemplates returns a JobTemplates
func newJobTemplates(c *FlowV1alpha1Client, namespace string) *jobTemplates {
	return &jobTemplates{
		gentype.NewClientWithListAndApply[*flowv1alpha1.JobTemplate, *flowv1alpha1.JobTemplateList, *applyconfigurationflowv1alpha1.JobTemplateApplyConfiguration](
			"jobtemplates",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *flowv1alpha1.JobTemplate { return &flowv1alpha1.JobTemplate{} },
			func() *flowv1alpha1.JobTemplateList { return &flowv1alpha1.JobTemplateList{} },
		),
	}
}
