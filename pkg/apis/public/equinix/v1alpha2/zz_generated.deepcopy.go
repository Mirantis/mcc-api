//go:build !ignore_autogenerated
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

package v1alpha2

import (
	v1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/public/equinix/v1alpha1"
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CephClusterConfig) DeepCopyInto(out *CephClusterConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CephClusterConfig.
func (in *CephClusterConfig) DeepCopy() *CephClusterConfig {
	if in == nil {
		return nil
	}
	out := new(CephClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EquinixMetalClusterProviderSpec) DeepCopyInto(out *EquinixMetalClusterProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.ClusterSpecMixin.DeepCopyInto(&out.ClusterSpecMixin)
	out.Ceph = in.Ceph
	in.BGP.DeepCopyInto(&out.BGP)
	in.Network.DeepCopyInto(&out.Network)
	if in.ProjectSSHKeys != nil {
		in, out := &in.ProjectSSHKeys, &out.ProjectSSHKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EquinixMetalClusterProviderSpec.
func (in *EquinixMetalClusterProviderSpec) DeepCopy() *EquinixMetalClusterProviderSpec {
	if in == nil {
		return nil
	}
	out := new(EquinixMetalClusterProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EquinixMetalClusterProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EquinixMetalClusterProviderStatus) DeepCopyInto(out *EquinixMetalClusterProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.ClusterStatusMixin.DeepCopyInto(&out.ClusterStatusMixin)
	if in.ElasticIPBlocks != nil {
		in, out := &in.ElasticIPBlocks, &out.ElasticIPBlocks
		*out = make([]v1alpha1.ElasticIPReservation, len(*in))
		copy(*out, *in)
	}
	if in.MetalLBRanges != nil {
		in, out := &in.MetalLBRanges, &out.MetalLBRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VlanStatus != nil {
		in, out := &in.VlanStatus, &out.VlanStatus
		*out = make(map[string]VlanStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EquinixMetalClusterProviderStatus.
func (in *EquinixMetalClusterProviderStatus) DeepCopy() *EquinixMetalClusterProviderStatus {
	if in == nil {
		return nil
	}
	out := new(EquinixMetalClusterProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EquinixMetalClusterProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EquinixMetalMachineProviderSpec) DeepCopyInto(out *EquinixMetalMachineProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.MachineSpecMixin.DeepCopyInto(&out.MachineSpecMixin)
	out.Ceph = in.Ceph
	if in.ProviderID != nil {
		in, out := &in.ProviderID, &out.ProviderID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(v1alpha1.Tags, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EquinixMetalMachineProviderSpec.
func (in *EquinixMetalMachineProviderSpec) DeepCopy() *EquinixMetalMachineProviderSpec {
	if in == nil {
		return nil
	}
	out := new(EquinixMetalMachineProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EquinixMetalMachineProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EquinixMetalMachineProviderStatus) DeepCopyInto(out *EquinixMetalMachineProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.MachineStatusMixin.DeepCopyInto(&out.MachineStatusMixin)
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]v1.NodeAddress, len(*in))
		copy(*out, *in)
	}
	if in.BGPNeighbors != nil {
		in, out := &in.BGPNeighbors, &out.BGPNeighbors
		*out = make([]v1alpha1.BGPNeighborPeers, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PrivateNetworks != nil {
		in, out := &in.PrivateNetworks, &out.PrivateNetworks
		*out = make([]v1alpha1.PrivateNetwork, len(*in))
		copy(*out, *in)
	}
	if in.VLANAddresses != nil {
		in, out := &in.VLANAddresses, &out.VLANAddresses
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EquinixMetalMachineProviderStatus.
func (in *EquinixMetalMachineProviderStatus) DeepCopy() *EquinixMetalMachineProviderStatus {
	if in == nil {
		return nil
	}
	out := new(EquinixMetalMachineProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EquinixMetalMachineProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAMConfig) DeepCopyInto(out *IPAMConfig) {
	*out = *in
	if in.IncludeRanges != nil {
		in, out := &in.IncludeRanges, &out.IncludeRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludeRanges != nil {
		in, out := &in.ExcludeRanges, &out.ExcludeRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Nameservers != nil {
		in, out := &in.Nameservers, &out.Nameservers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAMConfig.
func (in *IPAMConfig) DeepCopy() *IPAMConfig {
	if in == nil {
		return nil
	}
	out := new(IPAMConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
	if in.IncludeRanges != nil {
		in, out := &in.IncludeRanges, &out.IncludeRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludeRanges != nil {
		in, out := &in.ExcludeRanges, &out.ExcludeRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Nameservers != nil {
		in, out := &in.Nameservers, &out.Nameservers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DHCPRanges != nil {
		in, out := &in.DHCPRanges, &out.DHCPRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AdditionalVlans != nil {
		in, out := &in.AdditionalVlans, &out.AdditionalVlans
		*out = make([]Vlan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MetallbRanges != nil {
		in, out := &in.MetallbRanges, &out.MetallbRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network.
func (in *Network) DeepCopy() *Network {
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkStatus) DeepCopyInto(out *NetworkStatus) {
	*out = *in
	if in.IPAddrs != nil {
		in, out := &in.IPAddrs, &out.IPAddrs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkStatus.
func (in *NetworkStatus) DeepCopy() *NetworkStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vlan) DeepCopyInto(out *Vlan) {
	*out = *in
	if in.IncludeRanges != nil {
		in, out := &in.IncludeRanges, &out.IncludeRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludeRanges != nil {
		in, out := &in.ExcludeRanges, &out.ExcludeRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vlan.
func (in *Vlan) DeepCopy() *Vlan {
	if in == nil {
		return nil
	}
	out := new(Vlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VlanStatus) DeepCopyInto(out *VlanStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VlanStatus.
func (in *VlanStatus) DeepCopy() *VlanStatus {
	if in == nil {
		return nil
	}
	out := new(VlanStatus)
	in.DeepCopyInto(out)
	return out
}
