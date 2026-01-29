package stringset

import (
	"fmt"
	"strings"
)

type Set map[string]bool

func New() Set {
	return make(Set)
}

func NewFromSlice(l []string) Set {
	set := New()

	for _, s := range l {
		set.Add(s)
	}

	return set
}

func (s Set) String() string {
	keys := []string{}

	for key := range s {
		keys = append(keys, fmt.Sprintf("\"%s\"", key))
	}

	return fmt.Sprintf("{%s}", strings.Join(keys, ", "))
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	return s[elem]
}

func (s Set) Add(elem string) {
	s[elem] = true
}

func Subset(s1, s2 Set) bool {
	for key := range s1 {
		if !s2.Has(key) {
			return false
		}
	}

	return true
}

func Disjoint(s1, s2 Set) bool {
	return len(Intersection(s1, s2)) == 0
}

func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}

	for key := range s1 {
		if !s2.Has(key) {
			return false
		}
	}

	return true
}

func Intersection(s1, s2 Set) Set {
	intersected := New()

	for key := range s1 {
		if s2.Has(key) {
			intersected.Add(key)
		}
	}

	return intersected
}

func Difference(s1, s2 Set) Set {
	differences := New()

	for key := range s1 {
		if !s2.Has(key) {
			differences.Add(key)
		}
	}

	return differences
}

func Union(s1, s2 Set) Set {
	union := Set{}

	for key := range s1 {
		union[key] = true
	}

	for key := range s2 {
		union[key] = true
	}

	return union
}
