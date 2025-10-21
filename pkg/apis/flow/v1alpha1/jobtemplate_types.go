/*
Copyright 2021.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"volcano.sh/apis/pkg/apis/batch/v1alpha1"
)

// JobTemplateSpec defines the desired state of JobTemplate
type JobTemplateSpec struct {
	// JobSpec is the specification of the Job
	v1alpha1.JobSpec `json:"jobSpec,omitempty"`
}

// JobTemplateStatus defines the observed state of JobTemplate
type JobTemplateStatus struct {
	// JobDependsOnList is the list of jobs that this job depends on
	JobDependsOnList []string `json:"jobDependsOnList,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+kubebuilder:object:root=true
// +kubebuilder:resource:path=jobtemplates,shortName=jt
//+kubebuilder:subresource:status

// JobTemplate is the Schema for the jobtemplates API
type JobTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   v1alpha1.JobSpec  `json:"spec,omitempty"`
	Status JobTemplateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+kubebuilder:object:root=true

// JobTemplateList contains a list of JobTemplate
type JobTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []JobTemplate `json:"items"`
}
