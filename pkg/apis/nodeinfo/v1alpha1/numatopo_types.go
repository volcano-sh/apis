/*
Copyright 2020 The Kubernetes Authors.
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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ResourceInfo is the sets about resource capacity and allocatable
type ResourceInfo struct {
	Allocatable string `json:"allocatable,omitempty"`
	Capacity    int    `json:"capacity,omitempty"`
}

// CPUInfo is the cpu topology detail
type CPUInfo struct {
	NUMANodeID int `json:"numa,omitempty"`
	SocketID   int `json:"socket,omitempty"`
	CoreID     int `json:"core,omitempty"`
}

// MemoryTable contains memory information
type MemoryTable struct {
	TotalMemSize   uint64 `json:"total,omitempty"`
	SystemReserved uint64 `json:"systemReserved,omitempty"`
	Allocatable    uint64 `json:"allocatable,omitempty"`
	Reserved       uint64 `json:"reserved,omitempty"`
	Free           uint64 `json:"free,omitempty"`
}

// MemoryNode describes the node memory distribution
type MemoryNode struct {
	// NumberOfAssignments contains a number memory assignments from this node
	// When the container requires memory and hugepages it will increase number of assignments by two
	NumberOfAssignments int `json:"numberOfAssignments"`
	// MemoryTable contains NUMA node memory related information
	MemoryTable map[v1.ResourceName]*MemoryTable `json:"memoryMap,omitempty"`
	// Cells contains the current NUMA node and all other nodes that are in a group with current NUMA node
	// This parameter indicates if the current node is used for the multiple NUMA node memory allocation
	// For example if some container has pinning 0,1,2, NUMA nodes 0,1,2 under the state will have
	// this parameter equals to [0, 1, 2]
	Cells []int `json:"cells,omitempty"`
}

// PolicyName is the policy name type
type PolicyName string

const (
	// CPUManagerPolicy shows cpu manager policy type
	CPUManagerPolicy PolicyName = "CPUManagerPolicy"
	// MemoryManagerPolicy represents memory manager policy type
	MemoryManagerPolicy PolicyName = "MemoryManagerPolicy"
	// TopologyManagerPolicy shows topology manager policy type
	TopologyManagerPolicy PolicyName = "TopologyManagerPolicy"
)

// NumatopoSpec defines the desired state of NumaTopology
type NumatopoSpec struct {
	// Specifies the policy of the manager
	// +optional
	Policies map[PolicyName]string `json:"policies,omitempty"`

	// Specifies the reserved resource of the node
	// Key is resource name
	// +optional
	ResReserved map[string]string `json:"resReserved,omitempty"`

	// Specifies the numa info for the resource
	// Key is resource name
	// +optional
	NumaResMap map[string]ResourceInfo `json:"numares,omitempty"`

	// Specifies the cpu topology info
	// Key is cpu id
	// +optional
	CPUDetail map[string]CPUInfo `json:"cpuDetail,omitempty"`

	// Represents the memory topology info
	// mostly same as k8s/NUMANodeMap
	// key is numa node id
	// +optional
	MemoryDetail map[int]*MemoryNode `json:"MemoryDetail,omitempty"`
}

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=numatopo,scope=Cluster
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Numatopology is the Schema for the Numatopologies API
type Numatopology struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Specification of the numa information of the worker node
	Spec NumatopoSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NumatopologyList contains a list of Numatopology
type NumatopologyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Numatopology `json:"items"`
}
