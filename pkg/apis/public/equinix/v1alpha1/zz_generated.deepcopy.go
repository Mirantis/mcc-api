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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPConfig) DeepCopyInto(out *BGPConfig) {
	*out = *in
	if in.BIRDPeers != nil {
		in, out := &in.BIRDPeers, &out.BIRDPeers
		*out = make([]BGPPeer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MetalLBPeers != nil {
		in, out := &in.MetalLBPeers, &out.MetalLBPeers
		*out = make([]BGPPeer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPConfig.
func (in *BGPConfig) DeepCopy() *BGPConfig {
	if in == nil {
		return nil
	}
	out := new(BGPConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPNeighborPeers) DeepCopyInto(out *BGPNeighborPeers) {
	*out = *in
	if in.PeerIPs != nil {
		in, out := &in.PeerIPs, &out.PeerIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RoutesIn != nil {
		in, out := &in.RoutesIn, &out.RoutesIn
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RoutesOut != nil {
		in, out := &in.RoutesOut, &out.RoutesOut
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPNeighborPeers.
func (in *BGPNeighborPeers) DeepCopy() *BGPNeighborPeers {
	if in == nil {
		return nil
	}
	out := new(BGPNeighborPeers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPPeer) DeepCopyInto(out *BGPPeer) {
	*out = *in
	if in.PeerIPs != nil {
		in, out := &in.PeerIPs, &out.PeerIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPPeer.
func (in *BGPPeer) DeepCopy() *BGPPeer {
	if in == nil {
		return nil
	}
	out := new(BGPPeer)
	in.DeepCopyInto(out)
	return out
}

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
func (in *CephMachineConfig) DeepCopyInto(out *CephMachineConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CephMachineConfig.
func (in *CephMachineConfig) DeepCopy() *CephMachineConfig {
	if in == nil {
		return nil
	}
	out := new(CephMachineConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticIPReservation) DeepCopyInto(out *ElasticIPReservation) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticIPReservation.
func (in *ElasticIPReservation) DeepCopy() *ElasticIPReservation {
	if in == nil {
		return nil
	}
	out := new(ElasticIPReservation)
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
		*out = make([]ElasticIPReservation, len(*in))
		copy(*out, *in)
	}
	if in.MetalLBRanges != nil {
		in, out := &in.MetalLBRanges, &out.MetalLBRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
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
		*out = make(Tags, len(*in))
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
		*out = make([]BGPNeighborPeers, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PrivateNetworks != nil {
		in, out := &in.PrivateNetworks, &out.PrivateNetworks
		*out = make([]PrivateNetwork, len(*in))
		copy(*out, *in)
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
func (in *PrivateNetwork) DeepCopyInto(out *PrivateNetwork) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateNetwork.
func (in *PrivateNetwork) DeepCopy() *PrivateNetwork {
	if in == nil {
		return nil
	}
	out := new(PrivateNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Tags) DeepCopyInto(out *Tags) {
	{
		in := &in
		*out = make(Tags, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tags.
func (in Tags) DeepCopy() Tags {
	if in == nil {
		return nil
	}
	out := new(Tags)
	in.DeepCopyInto(out)
	return *out
}
