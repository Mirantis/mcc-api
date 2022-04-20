/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"crypto/rand"
	"time"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/wait"

	clusterv1 "github.com/Mirantis/mcc-api/pkg/apis/public/cluster/v1alpha1"
)

const (
	// MachineListFormatDeprecationMessage notifies the user that the old
	// MachineList format is no longer supported
	MachineListFormatDeprecationMessage = "Your MachineList items must include Kind and APIVersion"
)

// IsControlPlaneMachine checks machine is a control plane node.
func IsControlPlaneMachine(machine *clusterv1.Machine) bool {
	return machine.ObjectMeta.Labels[clusterv1.MachineControlPlaneLabelName] != ""
}

// IsBastionMachine checks is a bastion
func IsBastionMachine(machine *clusterv1.Machine) bool {
	return machine.ObjectMeta.Labels[clusterv1.MachineBastionLabelName] != ""
}

// Filter filters a list for a string.
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
func Contains(list []string, strToSearch string) bool {
	for _, item := range list {
		if item == strToSearch {
			return true
		}
	}
	return false
}

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
func PollImmediate(isInfinite bool, interval, timeout time.Duration, condition wait.ConditionFunc) error {
	if isInfinite {
		return wait.PollImmediateInfinite(interval, condition)
	}
	return wait.PollImmediate(interval, timeout, condition)
}
