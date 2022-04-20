/*
Copyright © 2021 Mirantis

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

package types

import "sort"

type IfAddressPlanItem struct {
	IfName    string
	Addresses []string
}
type IfAddressPlan []IfAddressPlanItem

func (r IfAddressPlan) Sort() {
	sort.Sort(r)
}

// Len for https://godoc.org/sort#Interface
func (r IfAddressPlan) Len() int {
	return len(r)
}

// Less for https://godoc.org/sort#Interface
func (r IfAddressPlan) Less(i, j int) bool {
	return r[i].IfName < r[j].IfName
}

// Swap for https://godoc.org/sort#Interface
func (r IfAddressPlan) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

type IfAddressPlanner interface {
	GetIfAddressPlan() IfAddressPlan
}

//-----------------------------------------------------------------------------

type Services2IfAddressPlanItem struct {
	IfName    string `json:"ifName,omitempty"`
	IPaddress string `json:"ipAddress,omitempty"`
}

type Services2IfAddressPlanList []Services2IfAddressPlanItem

func (r Services2IfAddressPlanList) Sort() {
	sort.Sort(r)
}

// Len for https://godoc.org/sort#Interface
func (r Services2IfAddressPlanList) Len() int {
	return len(r)
}

// Less for https://godoc.org/sort#Interface
func (r Services2IfAddressPlanList) Less(i, j int) bool {
	return r[i].IPaddress < r[j].IPaddress
}

// Swap for https://godoc.org/sort#Interface
func (r Services2IfAddressPlanList) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

type Services2IfAddressPlan map[string]Services2IfAddressPlanList
