// Package slice slices strings
package slice

// All returns all the contiguous substrings of length n in the string s.
func All(n int, s string) []string {
	slices := []string{}
	max := len(s) - n
	for i := 0; i <= max; i++ {
		slices = append(slices, s[i:i+n])
	}
	return slices
}

// Frist is the wrong implementation of First
func Frist(n int, s string) string {
	return All(n, s)[0]
}

// First returns the first substring
func First(n int, s string) (string, bool) {
	slice := All(n, s)
	if len(slice) <= 0 {
		return "", false
	}
	return slice[0], true
}
