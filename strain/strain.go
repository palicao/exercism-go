// Package strain implements keep and discard methods for various types
package strain

type Ints []int
type Lists [][]int
type Strings []string

// Keep for Ints
func (i Ints) Keep(f func(int) bool) Ints {
	var filtered Ints
	for _, el := range i {
		if f(el) == true {
			filtered = append(filtered, el)
		}
	}
	return filtered
}

// Discard for Ints
func (i Ints) Discard(f func(int) bool) Ints {
	return i.Keep(func(el int) bool {
		return !f(el)
	})
}

// Keep for Lists
func (l Lists) Keep(f func([]int) bool) Lists {
	var filtered Lists
	for _, el := range l {
		if f(el) == true {
			filtered = append(filtered, el)
		}
	}
	return filtered
}

// Keep for Strings
func (s Strings) Keep(f func(string) bool) Strings {
	var filtered Strings
	for _, el := range s {
		if f(el) == true {
			filtered = append(filtered, el)
		}
	}
	return filtered
}
