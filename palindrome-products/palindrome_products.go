// Package palindrome detects palindrome products in a given range.
package palindrome

import (
	"errors"
	"strconv"
)

type Product struct {
	Product int // palindromic, of course
	// list of all possible two-factor factorizations of Product, within
	// given limits, in order
	Factorizations [][2]int
}

// Products finds the minimum and the maximum palindrome products given a range of factors
func Products(fmin, fmax int) (pmin, pmax Product, e error) {

	if fmin > fmax {
		return Product{}, Product{}, errors.New("fmin > fmax...")
	}

	results := make(map[int]Product)
	keys := make(map[int]bool, 0)
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			p := i * j
			if isPalindrome(p) {
				if keys[p] {
					tmp := results[p]
					tmp.Factorizations = append(tmp.Factorizations, [2]int{i, j})
					results[p] = tmp
				} else {
					results[p] = Product{p, [][2]int{{i, j}}}
					keys[p] = true
				}
			}
		}
	}

	if len(keys) == 0 {
		return Product{}, Product{}, errors.New("No palindromes...")
	}

	var min, max int
	for key := range keys {
		if key > max {
			max = key
		}
		if min == 0 || key < min {
			min = key
		}
	}
	return results[min], results[max], nil
}

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	l := len(s)
	half := l / 2
	even := l % 2
	return s[:half] == strReverse(s[half+even:])
}

func strReverse(s string) string {
	b := make([]byte, len(s))
	var j int = len(s) - 1
	for i := 0; i <= j; i++ {
		b[j-i] = s[i]
	}
	return string(b)
}
