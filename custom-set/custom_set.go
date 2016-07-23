// Package stringset implements a custom set
package stringset

import "strings"

const testVersion = 3

type Set map[string]bool

// New returns a new Set
func New() Set {
	return Set{}
}

// NewFromSlice returns a new Set from a slice of strings
func NewFromSlice(s []string) Set {
	content := make(map[string]bool)
	for _, el := range s {
		content[el] = true
	}
	return Set(content)
}

// Add adds an element to the set
func (s Set) Add(el string) {
	s[el] = true
}

// Delete deletes an element from the set
func (s Set) Delete(el string) {
	delete(s, el)
}

// Has checks if a Set contains an element
func (s Set) Has(el string) bool {
	return s[el] == true
}

// IsEmpty checks if the Set is empty
func (s Set) IsEmpty() bool {
	return s.Len() == 0
}

// Len returns the size of a Set
func (s Set) Len() int {
	return len(s)
}

// Slice returns the Set as a slice
func (s Set) Slice() []string {
	slice := make([]string, 0)
	for el := range s {
		slice = append(slice, el)
	}
	return slice
}

// String returns the string representation of a Set
func (s Set) String() string {
	pieces := make([]string, 0)
	for el := range s {
		pieces = append(pieces, "\""+el+"\"")

	}
	return "{" + strings.Join(pieces, ", ") + "}"
}

// Equal checks if 2 Sets are equals
func Equal(s1, s2 Set) bool {
	return SymmetricDifference(s1, s2).Len() == 0
}

// Subset returns s1 ⊆ s2
func Subset(s1, s2 Set) bool {
	return Difference(s1, s2).Len() == 0
}

// Disjoint returns true if s1 and s2 have no element in common
func Disjoint(s1, s2 Set) bool {
	return Intersection(s1, s2).Len() == 0
}

// Intersection returns s1 ∩ s2
func Intersection(s1, s2 Set) Set {
	s := make(map[string]bool)
	for el := range s1 {
		if s2[el] {
			s[el] = true
		}
	}
	return Set(s)
}

// Union returns s1 ∪ s2
func Union(s1, s2 Set) Set {
	for el := range s2 {
		s1[el] = true
	}
	return Set(s1)
}

// Difference returns s1 ∖ s2
func Difference(s1, s2 Set) Set {
	for el := range s2 {
		if s1[el] {
			delete(s1, el)
		}
	}
	return Set(s1)
}

// SymmetricDifference returns elements that are in s1 or s2 but not on both
func SymmetricDifference(s1, s2 Set) Set {
	s := make(map[string]bool)
	for el := range s1 {
		if !s2[el] {
			s[el] = true
		}
	}
	for el := range s2 {
		if !s1[el] {
			s[el] = true
		}
	}
	return Set(s)
}
