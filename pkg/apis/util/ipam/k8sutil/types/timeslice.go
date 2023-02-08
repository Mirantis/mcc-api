package types

import (
	"sort"
	"time"
)

// +gocode:public-api=true
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
// +gocode:public-api=true
type TimeIndexItem struct {
	Time   time.Time
	Index  int
	String string
}

// +gocode:public-api=true
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
