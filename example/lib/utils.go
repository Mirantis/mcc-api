package lib

import (
	"time"

	pkgutil "github.com/Mirantis/mcc-api/pkg/util"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/klog/v2"
)

func RetryOnError(isWaitInfinite bool, fn func() error) error {
	var err error
	if waitErr := pkgutil.PollImmediate(isWaitInfinite, 10*time.Second, 5*time.Minute, func() (bool, error) {
		err = fn()
		if err != nil {
			if k8serrors.IsConflict(err) {
				return false, err
			}
			if !k8serrors.IsAlreadyExists(err) {
				klog.Infof("Error: %s, will retry", err)
				return false, nil
			}
		}
		return true, nil
	}); waitErr != nil {
		return err
	}
	return nil
}
