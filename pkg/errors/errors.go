package errors

import (
	"fmt"
	"go.uber.org/multierr"
	"k8s.io/klog"
)

// Number of callers skipped in the logs for the current function
// +gocode:public-api=true
const callersSkipForCurrentFunction = 1

// ErrorCollector contains a collection of errors objects.
// The object is not threadsafe.
// +gocode:public-api=true
type ErrorCollector struct {
	err error
	// Description of the error collection.
	description string
	// Add description to all logging messages.
	// Not enabled by default for compatibility with the old objects
	// because they may contain an uninformative description.
	useDescriptionForLogs bool
	// Number of callers skipped in the logs
	callerSkip int
}

// EnableDescriptionInLogs sets the property to add an additional description to the logs.
func (ec *ErrorCollector) EnableDescriptionInLogs() *ErrorCollector {
	ec.useDescriptionForLogs = true
	return ec
}

// DisableDescriptionInLogs sets the property to exclude an additional description from the logs
func (ec *ErrorCollector) DisableDescriptionInLogs() *ErrorCollector {
	ec.useDescriptionForLogs = false
	return ec
}

// AddCallerSkip increases the number of callers skipped in the logs
func (ec *ErrorCollector) AddCallerSkip(skip int) *ErrorCollector {
	ec.callerSkip += skip
	return ec
}

// GoString is an implementation of the fmt.GoStringer interface from stdlib:
// https://pkg.go.dev/fmt#GoStringer
func (ec *ErrorCollector) GoString() string {
	if ec == nil {
		return "(&ErrorCollector)(nil)"
	}

	errorsList := Errors(ec.err)
	errorsDesc := make([]string, 0, len(errorsList))
	for _, e := range errorsList {
		errorsDesc = append(errorsDesc, fmt.Sprintf("%#v", e))
	}

	return fmt.Sprintf(
		"&ErrorCollector{description: %s, err: %+v}",
		ec.description,
		errorsDesc,
	)
}

// Error is an implementation of the error interface from stdlib:
// https://pkg.go.dev/builtin#error
func (ec *ErrorCollector) Error() string {
	if ec != nil && ec.err != nil {
		return fmt.Sprintf(ec.description+": %s", ec.err)
	}

	return ""
}

// Collect adds an error to the collection with annotation
func (ec *ErrorCollector) Collect(err error, msg string) {
	ec.append(callersSkipForCurrentFunction, Wrap(err, msg))
}

// Collectf adds an error to the collection with annotation
func (ec *ErrorCollector) Collectf(err error, msg string, args ...interface{}) {
	ec.append(callersSkipForCurrentFunction, Wrapf(err, msg, args...))
}

// Append adds an error to the collection
func (ec *ErrorCollector) Append(err error) {
	ec.append(callersSkipForCurrentFunction, err)
}

// GetError returns an object which contains information about all collection errors.
func (ec *ErrorCollector) GetError() error {
	return ec.Unwrap()
}

// Is is an interface implementation for the errors.Is function from stdlib:
// https://pkg.go.dev/errors#Is
func (ec *ErrorCollector) Is(target error) bool {
	if ec == nil || ec.err == nil {
		return target == nil
	}

	return Is(ec.err, target)
}

// As is an interface implementation for the errors.As function from stdlib:
// https://pkg.go.dev/errors#As
func (ec *ErrorCollector) As(target interface{}) bool {
	if ec == nil || ec.err == nil {
		return false
	}

	return As(ec.err, target)
}

// Unwrap is an interface implementation for the errors.Unwrap function from stdlib:
// https://pkg.go.dev/errors#Unwrap
func (ec *ErrorCollector) Unwrap() error {
	if ec == nil {
		return nil
	}

	return ec.err
}
func (ec *ErrorCollector) append(depth int, err error) {
	if err == nil {
		return
	}

	ec.err = multierr.Append(ec.err, err)

	errMessage := err.Error()
	if ec.useDescriptionForLogs {
		errMessage = ec.description + ": " + errMessage
	}

	klog.InfoDepth(ec.callerSkip+depth+callersSkipForCurrentFunction, errMessage)
}

// +gocode:public-api=true
func NewErrorCollector(description string) *ErrorCollector {
	return &ErrorCollector{
		description: description,
	}
}
