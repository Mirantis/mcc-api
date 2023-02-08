package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// ProviderSpec defines the configuration to use during node creation.
// +gocode:public-api=true
type ProviderSpec struct {

	// Value is an inlined, serialized representation of the resource
	// configuration. It is recommended that providers maintain their own
	// versioned API types that should be serialized/deserialized from this
	// field, akin to component config.
	// +optional
	// +kubebuilder:validation:XPreserveUnknownFields
	Value *runtime.RawExtension `json:"value,omitempty"`

	// Source for the provider configuration. Cannot be used if value is
	// not empty.
	// +optional
	ValueFrom *ProviderSpecSource `json:"valueFrom,omitempty"`
}

// ProviderSpecSource represents a source for the provider-specific
// resource configuration.
type ProviderSpecSource struct {
	// The machine class from which the provider config should be sourced.
	// +optional
	MachineClass *MachineClassRef `json:"machineClass,omitempty"`
}

// MachineClassRef is a reference to the MachineClass object. Controllers should find the right MachineClass using this reference.
type MachineClassRef struct {
	// +optional
	*corev1.ObjectReference `json:",inline"`

	// Provider is the name of the cloud-provider which MachineClass is intended for.
	// +optional
	Provider string `json:"provider,omitempty"`
}
