// Package cryptosquare implements a simple square code encryption algorithm
package cryptosquare

import (
	"math"
	"strings"
)

const testVersion = 2

// Encodes encodes a string using the square code encryption algorithm
func Encode(s string) string {
	s = normalize(s)
	length := len(s)
	l := math.Sqrt(float64(length))
	c := int(math.Ceil(l))
	cols := make([]string, c)
	for i := 0; i < length; i++ {
		cols[i%c] += s[i : i+1]
	}
	return strings.Join(cols, " ")
}

// normalize converts to lowercase and removes non-alphanumeric characters
func normalize(s string) string {
	s = strings.ToLower(s)
	// tried with regexp, but it's much slower!
	// re := regexp.MustCompile("[^a-z0-9]")
	// return re.ReplaceAllString(s, "")
	normalized := ""
	for i := 0; i < len(s); i++ {
		l := s[i : i+1]
		if (l >= "a" && l <= "z") || (l >= "0" && l <= "9") {
			normalized += l
		}
	}
	return normalized
}
