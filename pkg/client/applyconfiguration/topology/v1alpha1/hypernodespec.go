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

// HyperNodeSpecApplyConfiguration represents a declarative configuration of the HyperNodeSpec type for use
// with apply.
type HyperNodeSpecApplyConfiguration struct {
	Tier    *int                           `json:"tier,omitempty"`
	Members []MemberSpecApplyConfiguration `json:"members,omitempty"`
}

// HyperNodeSpecApplyConfiguration constructs a declarative configuration of the HyperNodeSpec type for use with
// apply.
func HyperNodeSpec() *HyperNodeSpecApplyConfiguration {
	return &HyperNodeSpecApplyConfiguration{}
}

// WithTier sets the Tier field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Tier field is set to the value of the last call.
func (b *HyperNodeSpecApplyConfiguration) WithTier(value int) *HyperNodeSpecApplyConfiguration {
	b.Tier = &value
	return b
}

// WithMembers adds the given value to the Members field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Members field.
func (b *HyperNodeSpecApplyConfiguration) WithMembers(values ...*MemberSpecApplyConfiguration) *HyperNodeSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithMembers")
		}
		b.Members = append(b.Members, *values[i])
	}
	return b
}
