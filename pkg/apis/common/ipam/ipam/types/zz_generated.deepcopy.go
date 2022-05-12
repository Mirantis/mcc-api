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
func (in IfAddressPlan) DeepCopyInto(out *IfAddressPlan) {
	{
		in := &in
		*out = make(IfAddressPlan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IfAddressPlan.
func (in IfAddressPlan) DeepCopy() IfAddressPlan {
	if in == nil {
		return nil
	}
	out := new(IfAddressPlan)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IfAddressPlanItem) DeepCopyInto(out *IfAddressPlanItem) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IfAddressPlanItem.
func (in *IfAddressPlanItem) DeepCopy() *IfAddressPlanItem {
	if in == nil {
		return nil
	}
	out := new(IfAddressPlanItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Services2IfAddressPlan) DeepCopyInto(out *Services2IfAddressPlan) {
	{
		in := &in
		*out = make(Services2IfAddressPlan, len(*in))
		for key, val := range *in {
			var outVal []Services2IfAddressPlanItem
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(Services2IfAddressPlanList, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Services2IfAddressPlan.
func (in Services2IfAddressPlan) DeepCopy() Services2IfAddressPlan {
	if in == nil {
		return nil
	}
	out := new(Services2IfAddressPlan)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Services2IfAddressPlanItem) DeepCopyInto(out *Services2IfAddressPlanItem) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Services2IfAddressPlanItem.
func (in *Services2IfAddressPlanItem) DeepCopy() *Services2IfAddressPlanItem {
	if in == nil {
		return nil
	}
	out := new(Services2IfAddressPlanItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Services2IfAddressPlanList) DeepCopyInto(out *Services2IfAddressPlanList) {
	{
		in := &in
		*out = make(Services2IfAddressPlanList, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Services2IfAddressPlanList.
func (in Services2IfAddressPlanList) DeepCopy() Services2IfAddressPlanList {
	if in == nil {
		return nil
	}
	out := new(Services2IfAddressPlanList)
	in.DeepCopyInto(out)
	return *out
}