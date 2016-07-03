// Package hamming allows you to calculate the Hamming difference between two DNA strands.
package hamming

import "errors"

const testVersion = 4

// Distance returns Hamming distance between 2 DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("The size of the 2 strands must be the same")
	}
	diff := 0
	for i := range a {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff, nil
}
