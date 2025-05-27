package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Reservation defines the reservation.
// +kubebuilder:printcolumn:name="STATUS",type=string,JSONPath=`.status.state.phase`
// +kubebuilder:printcolumn:name="AGE",type=date,JSONPath=`.metadata.creationTimestamp`
// +kubebuilder:printcolumn:name="QUEUE",type=string,priority=1,JSONPath=`.spec.queue`
type Reservation struct {
	metav1.TypeMeta `json:",inline"`

	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Specification of the desired behavior of the reservation.
	// +optional
	Spec ReservationSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// Current status of the reservation.
	// +optional
	Status ReservationStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// ReservationSpec describes the desired behavior of the reservation.
type ReservationSpec struct {
	// SchedulerName is the default value of `tasks.template.spec.schedulerName`.
	// +optional
	SchedulerName string `json:"schedulerName,omitempty" protobuf:"bytes,1,opt,name=schedulerName"`

	// The minimal available pods to run for this Reservation
	// Defaults to the summary of tasks' replicas
	// +optional
	MinAvailable int32 `json:"minAvailable,omitempty" protobuf:"bytes,2,opt,name=minAvailable"`

	// Tasks specifies the task specification of Reservation
	// +optional
	Tasks []TaskSpec `json:"tasks,omitempty" protobuf:"bytes,3,opt,name=tasks"`

	//Specifies the queue that will be used in the scheduler, "default" queue is used this leaves empty.
	// +optional
	Queue string `json:"queue,omitempty" protobuf:"bytes,4,opt,name=queue"`

	// Owners specify the entities that can allocate the reserved resources.
	// Multiple owner selectors are ORed.
	Owners []ReservationOwner `json:"owners,omitempty" protobuf:"bytes,5,rep,name=owners"`

	// Time-to-Live period for the reservation.
	// `expires` and `ttl` are mutually exclusive.
	// +optional
	TTL *metav1.Duration `json:"ttl,omitempty" protobuf:"bytes,6,opt,name=ttl"`

	// Expired timestamp when the reservation is expected to expire.
	// If both `expires` and `ttl` are set, `expires` is checked first.
	// `expires` and `ttl` are mutually exclusive. Defaults to being set dynamically at runtime based on the `ttl`.
	// +optional
	Expires *metav1.Time `json:"expires,omitempty" protobuf:"bytes,7,opt,name=expires"`
}

// ReservationStatus represents the current status of a Reservation.
type ReservationStatus struct {
	// Current state of Reservation.
	// +optional
	State ReservationState `json:"state,omitempty" protobuf:"bytes,1,opt,name=state"`

	// The minimal available pods to run for this Reservation
	// +optional
	MinAvailable int32 `json:"minAvailable,omitempty" protobuf:"bytes,2,opt,name=minAvailable"`

	// The status of pods for each task
	// +optional
	TaskStatusCount map[string]TaskState `json:"taskStatusCount,omitempty" protobuf:"bytes,3,opt,name=taskStatusCount"`

	// The number of pending reservation pods.
	// +optional
	Pending int32 `json:"pending,omitempty" protobuf:"bytes,4,opt,name=pending"`

	// The number of available reservation pods.
	// +optional
	Available int32 `json:"available,omitempty" protobuf:"bytes,5,opt,name=available"`

	// The number of reservation pods which reached phase Succeeded.
	// +optional
	Succeeded int32 `json:"succeeded,omitempty" protobuf:"bytes,6,opt,name=succeeded"`

	// The number of reservation pods which reached phase Failed.
	// +optional
	Failed int32 `json:"failed,omitempty" protobuf:"bytes,7,opt,name=failed"`

	// Which conditions caused the current Reservation state.
	// +optional
	// +patchMergeKey=status
	// +patchStrategy=merge
	Conditions []ReservationCondition `json:"conditions,omitempty" protobuf:"bytes,8,rep,name=conditions"`

	// Owner who is currently using this reservation.
	// +optional
	CurrentOwner v1.ObjectReference `json:"currentOwner,omitempty" protobuf:"bytes,9,opt,name=currentOwner"`

	// Total allocatable resources for this reservation.
	// +optional
	Allocatable v1.ResourceList `json:"allocatable,omitempty" protobuf:"bytes,10,rep,name=allocatable"`

	// Total resources currently allocated to the owner.
	// +optional
	Allocated v1.ResourceList `json:"allocated,omitempty" protobuf:"bytes,11,rep,name=allocated"`
}

type ReservationState struct {
	// The phase of Reservation.
	// +optional
	Phase ReservationPhase `json:"phase,omitempty" protobuf:"bytes,1,opt,name=phase"`

	// Unique, one-word, CamelCase reason for the phase's last transition.
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,2,opt,name=reason"`

	// Human-readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,3,opt,name=message"`

	// Last time the condition transit from one phase to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,4,opt,name=lastTransitionTime"`
}

type ReservationCondition struct {
	// Status is the new phase of reservation after performing the state's action.
	Status ReservationPhase `json:"status" protobuf:"bytes,1,opt,name=status,casttype=ReservationPhase"`
	// Last time the condition transitioned from one phase to another.
	// +optional
	LastTransitionTime *metav1.Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,2,opt,name=lastTransitionTime"`
}

// ReservationPhase defines the phase of the reservation.
type ReservationPhase string

const (
	// ReservationPending is the phase that reservation is pending in the queue, waiting for scheduling decision
	ReservationPending ReservationPhase = "Pending"
	// ReservationAvailable indicates the Reservation is both scheduled and available for allocation.
	ReservationAvailable ReservationPhase = "Available"
	// ReservationSucceeded indicates the Reservation is scheduled and allocated for a owner, but not allocatable anymore.
	ReservationSucceeded ReservationPhase = "Succeeded"
	// ReservationWaiting indicates the Reservation is scheduled, but the resources to reserve are not ready for
	// allocation (e.g. in pre-allocation for running pods).
	ReservationWaiting ReservationPhase = "Waiting"
	// ReservationFailed indicates the Reservation is failed to reserve resources, due to expiration or marked as
	// unavailable, which the object is not available to allocate and will get cleaned in the future.
	ReservationFailed ReservationPhase = "Failed"
)

// ReservationOwner indicates the owner specification which can allocate reserved resources.
// +kubebuilder:validation:MinProperties=1
type ReservationOwner struct {
	// Multiple field selectors are ORed.
	// +optional
	Object *v1.ObjectReference `json:"object,omitempty" protobuf:"bytes,1,opt,name=object"`
	// +optional
	LabelSelector *metav1.LabelSelector `json:"labelSelector,omitempty" protobuf:"bytes,2,opt,name=labelSelector"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ReservationList defines the list of reservations.
type ReservationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []Reservation `json:"items" protobuf:"bytes,2,rep,name=items"`
}
