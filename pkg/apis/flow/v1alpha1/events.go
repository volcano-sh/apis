package v1alpha1

// Event represent the phase of JobFlow
// +kubebuilder:validation:Enum=OutOfSync
type Event string

const (
	// OutOfSyncEvent is triggered if JobFlow is updated(add/update/delete)
	OutOfSyncEvent Event = "OutOfSync"
)
