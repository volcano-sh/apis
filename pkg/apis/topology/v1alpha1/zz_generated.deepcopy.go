//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExactMatch) DeepCopyInto(out *ExactMatch) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExactMatch.
func (in *ExactMatch) DeepCopy() *ExactMatch {
	if in == nil {
		return nil
	}
	out := new(ExactMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperNode) DeepCopyInto(out *HyperNode) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperNode.
func (in *HyperNode) DeepCopy() *HyperNode {
	if in == nil {
		return nil
	}
	out := new(HyperNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HyperNode) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperNodeList) DeepCopyInto(out *HyperNodeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HyperNode, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperNodeList.
func (in *HyperNodeList) DeepCopy() *HyperNodeList {
	if in == nil {
		return nil
	}
	out := new(HyperNodeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HyperNodeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperNodeSpec) DeepCopyInto(out *HyperNodeSpec) {
	*out = *in
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]MemberSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperNodeSpec.
func (in *HyperNodeSpec) DeepCopy() *HyperNodeSpec {
	if in == nil {
		return nil
	}
	out := new(HyperNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperNodeStatus) DeepCopyInto(out *HyperNodeStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperNodeStatus.
func (in *HyperNodeStatus) DeepCopy() *HyperNodeStatus {
	if in == nil {
		return nil
	}
	out := new(HyperNodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberSelector) DeepCopyInto(out *MemberSelector) {
	*out = *in
	if in.ExactMatch != nil {
		in, out := &in.ExactMatch, &out.ExactMatch
		*out = new(ExactMatch)
		**out = **in
	}
	if in.RegexMatch != nil {
		in, out := &in.RegexMatch, &out.RegexMatch
		*out = new(RegexMatch)
		**out = **in
	}
	if in.LabelMatch != nil {
		in, out := &in.LabelMatch, &out.LabelMatch
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberSelector.
func (in *MemberSelector) DeepCopy() *MemberSelector {
	if in == nil {
		return nil
	}
	out := new(MemberSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberSpec) DeepCopyInto(out *MemberSpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberSpec.
func (in *MemberSpec) DeepCopy() *MemberSpec {
	if in == nil {
		return nil
	}
	out := new(MemberSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegexMatch) DeepCopyInto(out *RegexMatch) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegexMatch.
func (in *RegexMatch) DeepCopy() *RegexMatch {
	if in == nil {
		return nil
	}
	out := new(RegexMatch)
	in.DeepCopyInto(out)
	return out
}
