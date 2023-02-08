package v1alpha1

import (
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

type CustomJSON struct {
	apiextensions.JSON `json:",omitempty"`
}

func (in *CustomJSON) Set(val []byte) {
	in.JSON.Raw = val
}

func (in *CustomJSON) Get() []byte {
	return in.JSON.Raw
}
