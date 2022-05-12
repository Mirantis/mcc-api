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

package types

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in IfMapping) DeepCopyInto(out *IfMapping) {
	{
		in := &in
		*out = make(IfMapping, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IfMapping.
func (in IfMapping) DeepCopy() IfMapping {
	if in == nil {
		return nil
	}
	out := new(IfMapping)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in L3Layout) DeepCopyInto(out *L3Layout) {
	{
		in := &in
		*out = make(L3Layout, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L3Layout.
func (in L3Layout) DeepCopy() L3Layout {
	if in == nil {
		return nil
	}
	out := new(L3Layout)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L3SubnetSrc) DeepCopyInto(out *L3SubnetSrc) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = make(LabelSelector, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L3SubnetSrc.
func (in *L3SubnetSrc) DeepCopy() *L3SubnetSrc {
	if in == nil {
		return nil
	}
	out := new(L3SubnetSrc)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in LabelSelector) DeepCopyInto(out *LabelSelector) {
	{
		in := &in
		*out = make(LabelSelector, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LabelSelector.
func (in LabelSelector) DeepCopy() LabelSelector {
	if in == nil {
		return nil
	}
	out := new(LabelSelector)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NP) DeepCopyInto(out *NP) {
	*out = *in
	if in.l3Layout != nil {
		in, out := &in.l3Layout, &out.l3Layout
		*out = make(L3Layout, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Nics != nil {
		in, out := &in.Nics, &out.Nics
		*out = make(IfMapping, len(*in))
		copy(*out, *in)
	}
	if in.Macs != nil {
		in, out := &in.Macs, &out.Macs
		*out = make(IfMapping, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NP.
func (in *NP) DeepCopy() *NP {
	if in == nil {
		return nil
	}
	out := new(NP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Np4validation) DeepCopyInto(out *Np4validation) {
	*out = *in
	in.NP.DeepCopyInto(&out.NP)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Np4validation.
func (in *Np4validation) DeepCopy() *Np4validation {
	if in == nil {
		return nil
	}
	out := new(Np4validation)
	in.DeepCopyInto(out)
	return out
}