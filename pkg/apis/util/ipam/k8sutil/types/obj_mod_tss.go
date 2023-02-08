package types

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	apimTypes "k8s.io/apimachinery/pkg/types"
)

// +gocode:public-api=true
type ObjModTss interface {
	GetCreationTimestamp() v1.Time
	GetObjCreated() string
	GetObjUpdated() string
	GetObjStatusUpdated() string
	SetObjCreated(string) bool
	SetObjUpdated(string) bool
	SetObjStatusUpdated(string) bool
	GetNamespace() string
	GetName() string
	GetUID() apimTypes.UID
}
