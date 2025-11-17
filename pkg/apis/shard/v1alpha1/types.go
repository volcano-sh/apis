package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NodeShard is the Schema for the NodeShard API
type NodeShard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeShardSpec   `json:"spec,omitempty"`
	Status NodeShardStatus `json:"status,omitempty"`
}

// NodeShardSpec defines the desired state of NodeShard
type NodeShardSpec struct {
	// SchedulerName indicates the scheduler which is responsible for this NodeShard.
	// +optional
	SchedulerName string `json:"schedulerName,omitempty"`

	// NodesDesired contains the list of nodes that the scheduler wants in this shard.
	// +optional
	NodesDesired []string `json:"nodesDesired,omitempty"`
}

// NodeShardStatus defines the observed state of NodeShard
type NodeShardStatus struct {
	// LastUpdateTime is the last time the status was updated.
	// +optional
	LastUpdateTime *metav1.Time `json:"lastUpdateTime,omitempty"`

	// NodesInUse contains the list of nodes currently being used in this shard.
	// +optional
	NodesInUse []string `json:"nodesInUse,omitempty"`

	// NodesToRemove contains the list of nodes that are being removed from this shard.
	// +optional
	NodesToRemove []string `json:"nodesToRemove,omitempty"`

	// NodesToAdd contains the list of nodes that are being added to this shard.
	// +optional
	NodesToAdd []string `json:"nodesToAdd,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NodeShardList contains a list of NodeShard
type NodeShardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeShard `json:"items"`
}
