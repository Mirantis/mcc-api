package cidr32

import (
	"strings"
)

// IPRangeList -- list of IPRanges
// +gocode:public-api=true
type IPRangeList []IPRange

// Strings -- returns a list of string representation ranges
func (r IPRangeList) Strings() []string {
	rv := []string{}
	for _, r := range r {
		rv = append(rv, r.String())
	}
	return rv
}
func (r IPRangeList) String() string {
	return strings.Join(r.Strings(), "\n")
}

// ExcludeRange -- Remove exRange's addresses and returns pointer to new IPRangeList.
// also returns:
//   0 if no actions was
//   1 if amount of ranges unchanged
//   2 if amount of ranges was changed
func (r IPRangeList) ExcludeRange(exRange *IPRange) (rv IPRangeList, n int) {
	var tmp IPRangeList
	rv = append(IPRangeList{}, r...)
	e := len(rv)
	for i, shift := 0, 0; i < e; i++ {
		ner, nn := rv[i+shift].ExcludeRange(exRange)
		switch nn {
		case 0:

		case 1:

			rv[i+shift] = ner[0]
			if n == 0 {
				n = 1
			}
		case 2:

			if i+shift == 0 {
				tmp = append(IPRangeList{}, ner...)
				tmp = append(tmp, rv[1:]...)
			} else {
				tmp = append(IPRangeList{}, rv[:i+shift]...)
				tmp = append(tmp, ner[0])
				tmp = append(tmp, rv[i+shift:]...)
				tmp[i+shift+1] = ner[1]
			}
			shift++
			rv = tmp
			if n < 2 {
				n = 2
			}
		case -1:

			rv = append(rv[:i+shift], rv[i+shift+1:]...)
			e--
			if n < 2 {
				n = 2
			}
		}
	}
	return rv, n
}

// Capacity --
func (r IPRangeList) Capacity() (rv int) {
	for _, r := range r {
		rv += r.Capacity()
	}
	return rv
}

// Sorted --
func (r IPRangeList) Sorted() IPRangeList {

	return r
}

// Glued --
func (r IPRangeList) Glued() IPRangeList {

	return r
}

// Arranged -- sorted and Glued
func (r IPRangeList) Arranged() IPRangeList {

	tmp := r.Sorted()
	return tmp.Glued()
}
