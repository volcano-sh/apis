package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"volcano.sh/apis/pkg/apis/shard"
)

// SchemeGroupVersion is the group version used to register these objects.
var SchemeGroupVersion = schema.GroupVersion{Group: shard.GroupName, Version: "v1alpha1"}

// Resource takes an unqualified resource and returns a Group qualified Resource.
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &shard.SchemeBuilder
	// AddToScheme adds the types to the scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)
