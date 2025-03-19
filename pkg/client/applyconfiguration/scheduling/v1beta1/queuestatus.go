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

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	schedulingv1beta1 "volcano.sh/apis/pkg/apis/scheduling/v1beta1"
)

// QueueStatusApplyConfiguration represents a declarative configuration of the QueueStatus type for use
// with apply.
type QueueStatusApplyConfiguration struct {
	State       *schedulingv1beta1.QueueState  `json:"state,omitempty"`
	Unknown     *int32                         `json:"unknown,omitempty"`
	Pending     *int32                         `json:"pending,omitempty"`
	Running     *int32                         `json:"running,omitempty"`
	Inqueue     *int32                         `json:"inqueue,omitempty"`
	Completed   *int32                         `json:"completed,omitempty"`
	Reservation *ReservationApplyConfiguration `json:"reservation,omitempty"`
	Allocated   *v1.ResourceList               `json:"allocated,omitempty"`
}

// QueueStatusApplyConfiguration constructs a declarative configuration of the QueueStatus type for use with
// apply.
func QueueStatus() *QueueStatusApplyConfiguration {
	return &QueueStatusApplyConfiguration{}
}

// WithState sets the State field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the State field is set to the value of the last call.
func (b *QueueStatusApplyConfiguration) WithState(value schedulingv1beta1.QueueState) *QueueStatusApplyConfiguration {
	b.State = &value
	return b
}

// WithUnknown sets the Unknown field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Unknown field is set to the value of the last call.
func (b *QueueStatusApplyConfiguration) WithUnknown(value int32) *QueueStatusApplyConfiguration {
	b.Unknown = &value
	return b
}

// WithPending sets the Pending field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Pending field is set to the value of the last call.
func (b *QueueStatusApplyConfiguration) WithPending(value int32) *QueueStatusApplyConfiguration {
	b.Pending = &value
	return b
}

// WithRunning sets the Running field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Running field is set to the value of the last call.
func (b *QueueStatusApplyConfiguration) WithRunning(value int32) *QueueStatusApplyConfiguration {
	b.Running = &value
	return b
}

// WithInqueue sets the Inqueue field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Inqueue field is set to the value of the last call.
func (b *QueueStatusApplyConfiguration) WithInqueue(value int32) *QueueStatusApplyConfiguration {
	b.Inqueue = &value
	return b
}

// WithCompleted sets the Completed field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Completed field is set to the value of the last call.
func (b *QueueStatusApplyConfiguration) WithCompleted(value int32) *QueueStatusApplyConfiguration {
	b.Completed = &value
	return b
}

// WithReservation sets the Reservation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reservation field is set to the value of the last call.
func (b *QueueStatusApplyConfiguration) WithReservation(value *ReservationApplyConfiguration) *QueueStatusApplyConfiguration {
	b.Reservation = value
	return b
}

// WithAllocated sets the Allocated field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Allocated field is set to the value of the last call.
func (b *QueueStatusApplyConfiguration) WithAllocated(value v1.ResourceList) *QueueStatusApplyConfiguration {
	b.Allocated = &value
	return b
}
