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

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1alpha1 "volcano.sh/apis/v2/pkg/apis/batch/v1alpha1"
)

// JobRunningHistoryApplyConfiguration represents a declarative configuration of the JobRunningHistory type for use
// with apply.
type JobRunningHistoryApplyConfiguration struct {
	StartTimestamp *v1.Time           `json:"startTimestamp,omitempty"`
	EndTimestamp   *v1.Time           `json:"endTimestamp,omitempty"`
	State          *v1alpha1.JobPhase `json:"state,omitempty"`
}

// JobRunningHistoryApplyConfiguration constructs a declarative configuration of the JobRunningHistory type for use with
// apply.
func JobRunningHistory() *JobRunningHistoryApplyConfiguration {
	return &JobRunningHistoryApplyConfiguration{}
}

// WithStartTimestamp sets the StartTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StartTimestamp field is set to the value of the last call.
func (b *JobRunningHistoryApplyConfiguration) WithStartTimestamp(value v1.Time) *JobRunningHistoryApplyConfiguration {
	b.StartTimestamp = &value
	return b
}

// WithEndTimestamp sets the EndTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EndTimestamp field is set to the value of the last call.
func (b *JobRunningHistoryApplyConfiguration) WithEndTimestamp(value v1.Time) *JobRunningHistoryApplyConfiguration {
	b.EndTimestamp = &value
	return b
}

// WithState sets the State field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the State field is set to the value of the last call.
func (b *JobRunningHistoryApplyConfiguration) WithState(value v1alpha1.JobPhase) *JobRunningHistoryApplyConfiguration {
	b.State = &value
	return b
}
