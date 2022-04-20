package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const MCCUpgradeName = "mcc-upgrade"

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient
// +genclient:nonNamespaced

// MCCUpgrade configures upgrade schedule and provides status
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:subresource:status
type MCCUpgrade struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MCCUpgradeSpec   `json:"spec,omitempty"`
	Status MCCUpgradeStatus `json:"status,omitempty"`
}

type MCCUpgradeSpec struct {
	// Don't allow any upgrade operations before this time.
	BlockUntil *metav1.Time `json:"blockUntil,omitempty"`
	// Set of hours and weekdays when upgrade is allowed
	Schedule ScheduleItems `json:"schedule,omitempty"`
	// Timezone used for all schedule calculations, UTC if absent
	TimeZone string `json:"timeZone,omitempty"`
}

type ScheduleItems []ScheduleItem

type ScheduleItem struct {
	// Hours when upgrade is allowed
	Hours ScheduleHours `json:"hours,omitempty"`
	// Weekdays when upgrade is allowed
	Weekdays ScheduleWeekdays `json:"weekdays,omitempty"`
}

type ScheduleHours struct {
	// Block upgrades before this hour
	From int `json:"from,omitempty"`
	// Block upgrades after this hour
	To int `json:"to,omitempty"`
}

type ScheduleWeekdays struct {
	// Allow upgrades on Sunday
	Sunday bool `json:"sunday,omitempty"`
	// Allow upgrades on Monday
	Monday bool `json:"monday,omitempty"`
	// Allow upgrades on Tuesday
	Tuesday bool `json:"tuesday,omitempty"`
	// Allow upgrades on Wednesday
	Wednesday bool `json:"wednesday,omitempty"`
	// Allow upgrades on Thursday
	Thursday bool `json:"thursday,omitempty"`
	// Allow upgrades on Friday
	Friday bool `json:"friday,omitempty"`
	// Allow upgrades on Saturday
	Saturday bool `json:"saturday,omitempty"`
}

type MCCUpgradeStatus struct {
	NextRelease MCCUpgradeNextRelease `json:"nextRelease,omitempty"`
	NextAttempt metav1.Time           `json:"nextAttempt,omitempty"`
	Message     string                `json:"message,omitempty"`
}

type MCCUpgradeNextRelease struct {
	Version string      `json:"version,omitempty"`
	Date    metav1.Time `json:"date,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MCCUpgradeList is a list of MCCUpgrade objects
type MCCUpgradeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []MCCUpgrade `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MCCUpgrade{}, &MCCUpgradeList{})
}
