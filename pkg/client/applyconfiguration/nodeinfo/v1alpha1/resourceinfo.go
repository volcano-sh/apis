/*
Copyright 2021 The Volcano Authors.

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

// ResourceInfoApplyConfiguration represents an declarative configuration of the ResourceInfo type for use
// with apply.
type ResourceInfoApplyConfiguration struct {
	Allocatable *string `json:"allocatable,omitempty"`
	Capacity    *int    `json:"capacity,omitempty"`
}

// ResourceInfoApplyConfiguration constructs an declarative configuration of the ResourceInfo type for use with
// apply.
func ResourceInfo() *ResourceInfoApplyConfiguration {
	return &ResourceInfoApplyConfiguration{}
}

// WithAllocatable sets the Allocatable field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Allocatable field is set to the value of the last call.
func (b *ResourceInfoApplyConfiguration) WithAllocatable(value string) *ResourceInfoApplyConfiguration {
	b.Allocatable = &value
	return b
}

// WithCapacity sets the Capacity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Capacity field is set to the value of the last call.
func (b *ResourceInfoApplyConfiguration) WithCapacity(value int) *ResourceInfoApplyConfiguration {
	b.Capacity = &value
	return b
}
