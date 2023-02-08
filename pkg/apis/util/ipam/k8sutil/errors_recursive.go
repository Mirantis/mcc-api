package k8sutil

import (
	"errors"
)

// +gocode:public-api=true
func RecursiveUnwrap(err error) (rv error) {
	tmp := err
	for tmp != nil {
		rv = tmp
		tmp = errors.Unwrap(tmp)
	}
	return rv
}
