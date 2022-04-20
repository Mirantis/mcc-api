// +build !ignore_autogenerated

/*
Copyright 2022 The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BYOClusterProviderSpec) DeepCopyInto(out *BYOClusterProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.ClusterSpecMixin.DeepCopyInto(&out.ClusterSpecMixin)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BYOClusterProviderSpec.
func (in *BYOClusterProviderSpec) DeepCopy() *BYOClusterProviderSpec {
	if in == nil {
		return nil
	}
	out := new(BYOClusterProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BYOClusterProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BYOClusterProviderStatus) DeepCopyInto(out *BYOClusterProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.ClusterStatusMixin.DeepCopyInto(&out.ClusterStatusMixin)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BYOClusterProviderStatus.
func (in *BYOClusterProviderStatus) DeepCopy() *BYOClusterProviderStatus {
	if in == nil {
		return nil
	}
	out := new(BYOClusterProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BYOClusterProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BYOMachineProviderSpec) DeepCopyInto(out *BYOMachineProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.MachineSpecMixin.DeepCopyInto(&out.MachineSpecMixin)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BYOMachineProviderSpec.
func (in *BYOMachineProviderSpec) DeepCopy() *BYOMachineProviderSpec {
	if in == nil {
		return nil
	}
	out := new(BYOMachineProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BYOMachineProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BYOMachineProviderStatus) DeepCopyInto(out *BYOMachineProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.MachineStatusMixin.DeepCopyInto(&out.MachineStatusMixin)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BYOMachineProviderStatus.
func (in *BYOMachineProviderStatus) DeepCopy() *BYOMachineProviderStatus {
	if in == nil {
		return nil
	}
	out := new(BYOMachineProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BYOMachineProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
