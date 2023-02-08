package errors

import (
	stderrors "errors"
	"fmt"
	"go.uber.org/multierr"
)

// New is a function wrapper for stdlib/errors.New to avoid imports conflict.
// +gocode:public-api=true
func New(text string) error {
	return stderrors.New(text)
}

// Errorf is a function wrapper for stdlib/fmt.Errorf to avoid imports conflict.
// +gocode:public-api=true
func Errorf(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// As is a function wrapper for stdlib/errors.As to avoid imports conflict.
// +gocode:public-api=true
func As(err error, target interface{}) bool {
	if stderrors.As(err, target) {
		return true
	}

	if e := Cause(err); e != err && stderrors.As(e, target) {
		return true
	}

	return false
}

// Is is a function wrapper for stdlib/errors.Is to avoid imports conflict.
// +gocode:public-api=true
func Is(err, target error) bool {
	if stderrors.Is(err, target) {
		return true
	}

	if e := Cause(err); e != err && stderrors.Is(e, target) {
		return true
	}

	return false
}

// IsOneOf checks if the error is one of the targets
// +gocode:public-api=true
func IsOneOf(err error, target ...error) bool {
	for _, t := range target {
		if Is(err, t) {
			return true
		}
	}

	return false
}

// Unwrap is a function wrapper for stdlib/errors.Unwrap to avoid imports conflict.
// +gocode:public-api=true
func Unwrap(err error) error {
	return stderrors.Unwrap(err)
}

// Wrapf returns an error with annotation
// +gocode:public-api=true
func Wrapf(err error, msg string, args ...interface{}) error {
	if err != nil && msg != "" {
		description := fmt.Sprintf(msg, args...)
		return Errorf(description+": %w", err)
	}

	return err
}

// Wrap returns an error with annotation
// +gocode:public-api=true
func Wrap(err error, msg string) error {
	if err != nil && msg != "" {
		return Errorf(msg+": %w", err)
	}

	return err
}

// Errors returns a list of errors if it is a multiple error
// +gocode:public-api=true
func Errors(err error) []error {
	if err == nil {
		return nil
	}

	return multierr.Errors(err)
}

// Cause is a method for a compatibility with https://pkg.go.dev/github.com/pkg/errors#Cause
// Deprecated. Please use the method Unwrap.
// +gocode:public-api=true
func Cause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}

	return err
}
