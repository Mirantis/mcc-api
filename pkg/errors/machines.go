package errors

import (
	"fmt"
	commonerrors "github.com/Mirantis/mcc-api/v2/pkg/apis/cluster/common"
)

// A more descriptive kind of error that represents an error condition that
// should be set in the Machine.Status. The "Reason" field is meant for short,
// enum-style constants meant to be interpreted by machines. The "Message"
// field is meant to be read by humans.
// +gocode:public-api=true
type MachineError struct {
	Reason  commonerrors.MachineStatusError
	Message string
}

func (e *MachineError) Error() string {
	return e.Message
}

// Some error builders for ease of use. They set the appropriate "Reason"
// value, and all arguments are Printf-style varargs fed into Sprintf to
// construct the Message.
// +gocode:public-api=true
func InvalidMachineConfiguration(msg string, args ...interface{}) *MachineError {
	return &MachineError{
		Reason:  commonerrors.InvalidConfigurationMachineError,
		Message: fmt.Sprintf(msg, args...),
	}
}

// +gocode:public-api=true
func CreateMachine(msg string, args ...interface{}) *MachineError {
	return &MachineError{
		Reason:  commonerrors.CreateMachineError,
		Message: fmt.Sprintf(msg, args...),
	}
}

// +gocode:public-api=true
func UpdateMachine(msg string, args ...interface{}) *MachineError {
	return &MachineError{
		Reason:  commonerrors.UpdateMachineError,
		Message: fmt.Sprintf(msg, args...),
	}
}

// +gocode:public-api=true
func DeleteMachine(msg string, args ...interface{}) *MachineError {
	return &MachineError{
		Reason:  commonerrors.DeleteMachineError,
		Message: fmt.Sprintf(msg, args...),
	}
}
