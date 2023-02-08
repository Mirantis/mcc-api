package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Cluster stores info about the target cluster which must be tested
type Cluster struct {
	Name string `json:"name"`
}

// Pipeline defines data needed to get test execution logic
type Pipeline struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

// TestRun is the Schema for the testruns API
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=kaas
// +kubebuilder:printcolumn:name="Cluster Name",type="string",JSONPath=".spec.cluster.name"
// +kubebuilder:printcolumn:name="Succeeded",type="string",JSONPath=".status.succeeded"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.reason"
// +kubebuilder:printcolumn:name="StartTime",type="date",JSONPath=".status.startTime"
// +kubebuilder:printcolumn:name="CompletionTime",type="date",JSONPath=".status.completionTime"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message",priority=1
// +kubebuilder:printcolumn:name="Results locations",type="string",JSONPath=".status.resultsLocation",priority=1
// +gocode:public-api=true
type TestRun struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec TestRunSpec `json:"spec"`
	// +optional
	Status TestRunStatus `json:"status,omitempty"`
}

// TestRunStatus is status subresource for a TestRun resource
type TestRunStatus struct {
	// +optional
	StartTime *metav1.Time `json:"startTime,omitempty"`
	// +optional
	CompletionTime *metav1.Time `json:"completionTime,omitempty"`
	// +optional
	Succeeded corev1.ConditionStatus `json:"succeeded,omitempty"`
	// +optional
	Reason string `json:"reason,omitempty"`
	// +optional
	Message string `json:"message,omitempty"`
	// +optional
	ResultsLocation string `json:"resultsLocation,omitempty"`
	// +optional
	Tasks map[string]TaskStatus `json:"tasks,omitempty"`
}

// TestRunList is a list of TestRun resources
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type TestRunList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []TestRun `json:"items"`
}

// Test abstracts the implementation (Tekton) from the user interface
// +genclient
// +genclient:noStatus
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +gocode:public-api=true
type Test struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec TestSpec `json:"spec"`
}

// TestList is a list of Test resources
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type TestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Test `json:"items"`
}

// TestRunSpec is spec subresource for a TestRun resource
type TestRunSpec struct {
	Cluster Cluster `json:"cluster"`
	TestRef string  `json:"testRef"`
	Params  []Param `json:"params,omitempty"`
	// +optional
	Timeout *metav1.Duration `json:"timeout,omitempty"`
}

// TaskStatus is status subresource for a Tekton's TaskRun resource
type TaskStatus struct {
	// +optional
	StartTime *metav1.Time `json:"startTime,omitempty"`
	// +optional
	CompletionTime *metav1.Time `json:"completionTime,omitempty"`
	// +optional
	Succeeded corev1.ConditionStatus `json:"succeeded,omitempty"`
	// +optional
	Reason string `json:"reason,omitempty"`
	// +optional
	Message string `json:"message,omitempty"`
}

// Param declares a mapping of names and values
type Param struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// TestSpec defines the desired state of a Test
type TestSpec struct {
	// Pipeline stores the name and namespace of a Pipeline to be used by a TestRun
	Pipeline Pipeline `json:"pipeline"`
	// Params declares a list of input parameters that must be supplied
	Params []ParamSpec `json:"params"`
	// Description is a user-facing description of a Test CRD
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Timeout *metav1.Duration `json:"timeout,omitempty"`
}

// ParamSpec defines arbitrary parameters needed beyond typed inputs
type ParamSpec struct {
	// Name declares the name by which a parameter is referenced
	Name string `json:"name"`
	// Description stores user-facing info about param
	// +optional
	Description string `json:"description,omitempty"`
	// Default is the value a parameter takes if no input value is supplied
	// +optional
	Default string `json:"default,omitempty"`
}
