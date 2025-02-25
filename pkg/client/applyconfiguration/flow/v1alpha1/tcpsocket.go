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

// TcpSocketApplyConfiguration represents a declarative configuration of the TcpSocket type for use
// with apply.
type TcpSocketApplyConfiguration struct {
	TaskName *string `json:"taskName,omitempty"`
	Port     *int    `json:"port,omitempty"`
}

// TcpSocketApplyConfiguration constructs a declarative configuration of the TcpSocket type for use with
// apply.
func TcpSocket() *TcpSocketApplyConfiguration {
	return &TcpSocketApplyConfiguration{}
}

// WithTaskName sets the TaskName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TaskName field is set to the value of the last call.
func (b *TcpSocketApplyConfiguration) WithTaskName(value string) *TcpSocketApplyConfiguration {
	b.TaskName = &value
	return b
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *TcpSocketApplyConfiguration) WithPort(value int) *TcpSocketApplyConfiguration {
	b.Port = &value
	return b
}
