package types

import (
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// +gocode:public-api=true
type AutoUpdatableK8sObj interface {
	client.Object
	StatusToYAML() ([]byte, error)
	YAMLtoStatus([]byte) error
	GetStatus() interface{}
	GetMetadata() interface{}
	GetSpec() interface{}
	SetObjUpdated(string) bool
	SetObjStatusUpdated(string) bool
}
