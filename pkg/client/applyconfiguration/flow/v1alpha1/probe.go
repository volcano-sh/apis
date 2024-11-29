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

// ProbeApplyConfiguration represents a declarative configuration of the Probe type for use
// with apply.
type ProbeApplyConfiguration struct {
	HttpGetList    []HttpGetApplyConfiguration    `json:"httpGetList,omitempty"`
	TcpSocketList  []TcpSocketApplyConfiguration  `json:"tcpSocketList,omitempty"`
	TaskStatusList []TaskStatusApplyConfiguration `json:"taskStatusList,omitempty"`
}

// ProbeApplyConfiguration constructs a declarative configuration of the Probe type for use with
// apply.
func Probe() *ProbeApplyConfiguration {
	return &ProbeApplyConfiguration{}
}

// WithHttpGetList adds the given value to the HttpGetList field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the HttpGetList field.
func (b *ProbeApplyConfiguration) WithHttpGetList(values ...*HttpGetApplyConfiguration) *ProbeApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithHttpGetList")
		}
		b.HttpGetList = append(b.HttpGetList, *values[i])
	}
	return b
}

// WithTcpSocketList adds the given value to the TcpSocketList field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the TcpSocketList field.
func (b *ProbeApplyConfiguration) WithTcpSocketList(values ...*TcpSocketApplyConfiguration) *ProbeApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTcpSocketList")
		}
		b.TcpSocketList = append(b.TcpSocketList, *values[i])
	}
	return b
}

// WithTaskStatusList adds the given value to the TaskStatusList field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the TaskStatusList field.
func (b *ProbeApplyConfiguration) WithTaskStatusList(values ...*TaskStatusApplyConfiguration) *ProbeApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTaskStatusList")
		}
		b.TaskStatusList = append(b.TaskStatusList, *values[i])
	}
	return b
}