/*
Copyright © 2021 Mirantis

Inspired by https://github.com/inwinstack/ipam/, https://github.com/inwinstack/blended/

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

import (
	"sort"
	"time"
)

type TimeSlice []time.Time

func (r TimeSlice) Sort() {
	sort.Sort(r)
}

// Len for https://godoc.org/sort#Interface
func (r TimeSlice) Len() int {
	return len(r)
}

// Less for https://godoc.org/sort#Interface
func (r TimeSlice) Less(i, j int) bool {
	return r[i].Before(r[j])
}

// Swap for https://godoc.org/sort#Interface
func (r TimeSlice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

//-----------------------------------------------------------------------------

type TimeIndexItem struct {
	Time   time.Time
	Index  int
	String string
}
type TimeIndexSlice []TimeIndexItem

func (r TimeIndexSlice) Sort() {
	sort.Sort(r)
}

// Len for https://godoc.org/sort#Interface
func (r TimeIndexSlice) Len() int {
	return len(r)
}

// Less for https://godoc.org/sort#Interface
func (r TimeIndexSlice) Less(i, j int) bool {
	if r[i].Time.IsZero() {
		return false
	}
	if r[j].Time.IsZero() {
		return true
	}
	return r[i].Time.Before(r[j].Time)
}

// Swap for https://godoc.org/sort#Interface
func (r TimeIndexSlice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
