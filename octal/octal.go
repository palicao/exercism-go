// Package octal converts number from octal to decimal
package octal

import (
	"errors"
	"math"
)

// ParseOctal converts an octal number to decimal
func ParseOctal(o string) (d int64, e error) {
	l := float64(len(o) - 1)
	eight := float64(8)
	for _, i := range o {
		if !isOctal(i) {
			return 0, errors.New("The string is not valid octal")
		}
		d += int64(i-'0') * int64(math.Pow(eight, l))
		l--
	}
	return d, nil
}

func isOctal(s rune) bool {
	if s >= '0' && s <= '7' {
		return true
	}
	return false
}
