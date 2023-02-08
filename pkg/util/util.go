package util

import (
	"crypto/rand"
	clusterv1 "github.com/Mirantis/mcc-api/v2/pkg/apis/cluster/v1alpha1"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/wait"
	"time"
)

const (
	// MachineListFormatDeprecationMessage notifies the user that the old
	// MachineList format is no longer supported
	// +gocode:public-api=true
	MachineListFormatDeprecationMessage = "Your MachineList items must include Kind and APIVersion"
)

// IsControlPlaneMachine checks machine is a control plane node.
// +gocode:public-api=true
func IsControlPlaneMachine(machine *clusterv1.Machine) bool {
	return machine.ObjectMeta.Labels[clusterv1.MachineControlPlaneLabelName] != ""
}

// IsBastionMachine checks is a bastion
// +gocode:public-api=true
func IsBastionMachine(machine *clusterv1.Machine) bool {
	return machine.ObjectMeta.Labels[clusterv1.MachineBastionLabelName] != ""
}

// Filter filters a list for a string.
// +gocode:public-api=true
func Filter(list []string, strToFilter string) []string {
	var newList []string

	for _, item := range list {
		if item != strToFilter {
			newList = append(newList, item)
		}
	}

	return newList
}

// Contains returns true if a list contains a string.
// +gocode:public-api=true
func Contains(list []string, strToSearch string) bool {
	for _, item := range list {
		if item == strToSearch {
			return true
		}
	}
	return false
}

// +gocode:public-api=true
func GenerateRandomBytes(length int) (data []byte, err error) {
	data = make([]byte, length)
	var n int
	if n, err = rand.Read(data); err != nil {
		return nil, err
	} else if n != length {
		return nil, errors.Errorf(
			"Not enough data. Expected to read: %v bytes, got: %v bytes", length, n,
		)
	}
	return data, nil
}

// +gocode:public-api=true
func GenerateRandomString(length int) (string, error) {
	keyBytes, err := GenerateRandomBytes(length)
	if err != nil {
		return "", err
	}
	chars := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsLen := len(chars)
	for i := 0; i < length; i++ {
		keyBytes[i] = chars[keyBytes[i]%byte(charsLen)]
	}
	return string(keyBytes), nil
}

// Diff compares two lists and returns unique items from the source list
// which are not contained in the target list.
// +gocode:public-api=true
func Diff(target, source []string) []string {
	diff := make([]string, 0, len(source))
	for _, v := range source {
		if !Contains(diff, v) && !Contains(target, v) {
			diff = append(diff, v)
		}
	}

	return diff
}

// PollImmediate is warpper to call wait.PollImmediate or PollImmediateInfinite functions
// +gocode:public-api=true
func PollImmediate(isInfinite bool, interval, timeout time.Duration, condition wait.ConditionFunc) error {
	if isInfinite {
		return wait.PollImmediateInfinite(interval, condition)
	}
	return wait.PollImmediate(interval, timeout, condition)
}
