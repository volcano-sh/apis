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
)

// PodGroupSpecApplyConfiguration represents a declarative configuration of the PodGroupSpec type for use
// with apply.
type PodGroupSpecApplyConfiguration struct {
	MinMember         *int32                                 `json:"minMember,omitempty"`
	MinTaskMember     map[string]int32                       `json:"minTaskMember,omitempty"`
	Queue             *string                                `json:"queue,omitempty"`
	PriorityClassName *string                                `json:"priorityClassName,omitempty"`
	MinResources      *v1.ResourceList                       `json:"minResources,omitempty"`
	NetworkTopology   *NetworkTopologySpecApplyConfiguration `json:"networkTopology,omitempty"`
}

// PodGroupSpecApplyConfiguration constructs a declarative configuration of the PodGroupSpec type for use with
// apply.
func PodGroupSpec() *PodGroupSpecApplyConfiguration {
	return &PodGroupSpecApplyConfiguration{}
}

// WithMinMember sets the MinMember field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MinMember field is set to the value of the last call.
func (b *PodGroupSpecApplyConfiguration) WithMinMember(value int32) *PodGroupSpecApplyConfiguration {
	b.MinMember = &value
	return b
}

// WithMinTaskMember puts the entries into the MinTaskMember field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the MinTaskMember field,
// overwriting an existing map entries in MinTaskMember field with the same key.
func (b *PodGroupSpecApplyConfiguration) WithMinTaskMember(entries map[string]int32) *PodGroupSpecApplyConfiguration {
	if b.MinTaskMember == nil && len(entries) > 0 {
		b.MinTaskMember = make(map[string]int32, len(entries))
	}
	for k, v := range entries {
		b.MinTaskMember[k] = v
	}
	return b
}

// WithQueue sets the Queue field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Queue field is set to the value of the last call.
func (b *PodGroupSpecApplyConfiguration) WithQueue(value string) *PodGroupSpecApplyConfiguration {
	b.Queue = &value
	return b
}

// WithPriorityClassName sets the PriorityClassName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PriorityClassName field is set to the value of the last call.
func (b *PodGroupSpecApplyConfiguration) WithPriorityClassName(value string) *PodGroupSpecApplyConfiguration {
	b.PriorityClassName = &value
	return b
}

// WithMinResources sets the MinResources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MinResources field is set to the value of the last call.
func (b *PodGroupSpecApplyConfiguration) WithMinResources(value v1.ResourceList) *PodGroupSpecApplyConfiguration {
	b.MinResources = &value
	return b
}

// WithNetworkTopology sets the NetworkTopology field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NetworkTopology field is set to the value of the last call.
func (b *PodGroupSpecApplyConfiguration) WithNetworkTopology(value *NetworkTopologySpecApplyConfiguration) *PodGroupSpecApplyConfiguration {
	b.NetworkTopology = value
	return b
}
