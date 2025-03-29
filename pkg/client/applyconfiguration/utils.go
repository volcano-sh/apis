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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
	v1alpha1 "volcano.sh/apis/pkg/apis/batch/v1alpha1"
	busv1alpha1 "volcano.sh/apis/pkg/apis/bus/v1alpha1"
	flowv1alpha1 "volcano.sh/apis/pkg/apis/flow/v1alpha1"
	nodeinfov1alpha1 "volcano.sh/apis/pkg/apis/nodeinfo/v1alpha1"
	v1beta1 "volcano.sh/apis/pkg/apis/scheduling/v1beta1"
	batchv1alpha1 "volcano.sh/apis/pkg/client/applyconfiguration/batch/v1alpha1"
	applyconfigurationbusv1alpha1 "volcano.sh/apis/pkg/client/applyconfiguration/bus/v1alpha1"
	applyconfigurationflowv1alpha1 "volcano.sh/apis/pkg/client/applyconfiguration/flow/v1alpha1"
	internal "volcano.sh/apis/pkg/client/applyconfiguration/internal"
	applyconfigurationnodeinfov1alpha1 "volcano.sh/apis/pkg/client/applyconfiguration/nodeinfo/v1alpha1"
	schedulingv1beta1 "volcano.sh/apis/pkg/client/applyconfiguration/scheduling/v1beta1"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=batch.volcano.sh, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("DependsOn"):
		return &batchv1alpha1.DependsOnApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Job"):
		return &batchv1alpha1.JobApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("JobCondition"):
		return &batchv1alpha1.JobConditionApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("JobSpec"):
		return &batchv1alpha1.JobSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("JobState"):
		return &batchv1alpha1.JobStateApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("JobStatus"):
		return &batchv1alpha1.JobStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("LifecyclePolicy"):
		return &batchv1alpha1.LifecyclePolicyApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("TaskSpec"):
		return &batchv1alpha1.TaskSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("TaskState"):
		return &batchv1alpha1.TaskStateApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("VolumeSpec"):
		return &batchv1alpha1.VolumeSpecApplyConfiguration{}

		// Group=bus.volcano.sh, Version=v1alpha1
	case busv1alpha1.SchemeGroupVersion.WithKind("Command"):
		return &applyconfigurationbusv1alpha1.CommandApplyConfiguration{}

		// Group=flow.volcano.sh, Version=v1alpha1
	case flowv1alpha1.SchemeGroupVersion.WithKind("Condition"):
		return &applyconfigurationflowv1alpha1.ConditionApplyConfiguration{}
	case flowv1alpha1.SchemeGroupVersion.WithKind("DependsOn"):
		return &applyconfigurationflowv1alpha1.DependsOnApplyConfiguration{}
	case flowv1alpha1.SchemeGroupVersion.WithKind("Flow"):
		return &applyconfigurationflowv1alpha1.FlowApplyConfiguration{}
	case flowv1alpha1.SchemeGroupVersion.WithKind("HttpGet"):
		return &applyconfigurationflowv1alpha1.HttpGetApplyConfiguration{}
	case flowv1alpha1.SchemeGroupVersion.WithKind("JobFlow"):
		return &applyconfigurationflowv1alpha1.JobFlowApplyConfiguration{}
	case flowv1alpha1.SchemeGroupVersion.WithKind("JobFlowSpec"):
		return &applyconfigurationflowv1alpha1.JobFlowSpecApplyConfiguration{}
	case flowv1alpha1.SchemeGroupVersion.WithKind("JobFlowStatus"):
		return &applyconfigurationflowv1alpha1.JobFlowStatusApplyConfiguration{}
	case flowv1alpha1.SchemeGroupVersion.WithKind("JobRunningHistory"):
		return &applyconfigurationflowv1alpha1.JobRunningHistoryApplyConfiguration{}
	case flowv1alpha1.SchemeGroupVersion.WithKind("JobStatus"):
		return &applyconfigurationflowv1alpha1.JobStatusApplyConfiguration{}
	case flowv1alpha1.SchemeGroupVersion.WithKind("JobTemplate"):
		return &applyconfigurationflowv1alpha1.JobTemplateApplyConfiguration{}
	case flowv1alpha1.SchemeGroupVersion.WithKind("JobTemplateStatus"):
		return &applyconfigurationflowv1alpha1.JobTemplateStatusApplyConfiguration{}
	case flowv1alpha1.SchemeGroupVersion.WithKind("Probe"):
		return &applyconfigurationflowv1alpha1.ProbeApplyConfiguration{}
	case flowv1alpha1.SchemeGroupVersion.WithKind("State"):
		return &applyconfigurationflowv1alpha1.StateApplyConfiguration{}
	case flowv1alpha1.SchemeGroupVersion.WithKind("Target"):
		return &applyconfigurationflowv1alpha1.TargetApplyConfiguration{}
	case flowv1alpha1.SchemeGroupVersion.WithKind("TaskStatus"):
		return &applyconfigurationflowv1alpha1.TaskStatusApplyConfiguration{}
	case flowv1alpha1.SchemeGroupVersion.WithKind("TcpSocket"):
		return &applyconfigurationflowv1alpha1.TcpSocketApplyConfiguration{}

		// Group=nodeinfo.volcano.sh, Version=v1alpha1
	case nodeinfov1alpha1.SchemeGroupVersion.WithKind("CPUInfo"):
		return &applyconfigurationnodeinfov1alpha1.CPUInfoApplyConfiguration{}
	case nodeinfov1alpha1.SchemeGroupVersion.WithKind("Numatopology"):
		return &applyconfigurationnodeinfov1alpha1.NumatopologyApplyConfiguration{}
	case nodeinfov1alpha1.SchemeGroupVersion.WithKind("NumatopoSpec"):
		return &applyconfigurationnodeinfov1alpha1.NumatopoSpecApplyConfiguration{}
	case nodeinfov1alpha1.SchemeGroupVersion.WithKind("ResourceInfo"):
		return &applyconfigurationnodeinfov1alpha1.ResourceInfoApplyConfiguration{}

		// Group=scheduling.volcano.sh, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithKind("Affinity"):
		return &schedulingv1beta1.AffinityApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("Cluster"):
		return &schedulingv1beta1.ClusterApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("Guarantee"):
		return &schedulingv1beta1.GuaranteeApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("NodeGroupAffinity"):
		return &schedulingv1beta1.NodeGroupAffinityApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("NodeGroupAntiAffinity"):
		return &schedulingv1beta1.NodeGroupAntiAffinityApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("PodGroup"):
		return &schedulingv1beta1.PodGroupApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("PodGroupCondition"):
		return &schedulingv1beta1.PodGroupConditionApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("PodGroupSpec"):
		return &schedulingv1beta1.PodGroupSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("PodGroupStatus"):
		return &schedulingv1beta1.PodGroupStatusApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("Queue"):
		return &schedulingv1beta1.QueueApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("QueueSpec"):
		return &schedulingv1beta1.QueueSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("QueueStatus"):
		return &schedulingv1beta1.QueueStatusApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("Reservation"):
		return &schedulingv1beta1.ReservationApplyConfiguration{}

	}
	return nil
}

func NewTypeConverter(scheme *runtime.Scheme) *testing.TypeConverter {
	return &testing.TypeConverter{Scheme: scheme, TypeResolver: internal.Parser()}
}
