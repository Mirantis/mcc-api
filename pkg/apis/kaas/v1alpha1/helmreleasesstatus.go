package v1alpha1

import (
	crclient "sigs.k8s.io/controller-runtime/pkg/client"
	"strings"
)

// +gocode:public-api=true
type ComponentStatus struct {
	URL string `json:"url,omitempty"`
}

// +gocode:public-api=true
type HelmReleasesStatus struct {
	StackLight      StackLightStatus         `json:"stacklight,omitempty"`
	IAM             IAMStatus                `json:"iam,omitempty"`
	DECC            DECCStatus               `json:"decc,omitempty"`
	ReleaseStatuses map[string]ReleaseStatus `json:"releaseStatuses,omitempty"`
}

func (s HelmReleasesStatus) Error() string {
	message := "not ready helm releases: "
	for release, status := range s.ReleaseStatuses {
		if !status.Success {
			message += release + ", "
		}
	}
	return strings.TrimSuffix(message, ", ")
}

type ReleaseStatus struct {
	Success bool `json:"success,omitempty"`
}

// +gocode:public-api=true
type ComponentStatusGetter interface {
	GetStatus(client crclient.Client, serviceName string) ComponentStatus
}
