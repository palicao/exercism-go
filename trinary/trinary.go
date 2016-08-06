// Package trinary converts trinary numbers into decimal
package trinary

import "errors"

// ParseTrinary converts a trinary number to decimal
func ParseTrinary(trinary string) (decimal int64, err error) {
	len := len(trinary)
	for i, t := range trinary {
		if t < '0' || t > '2' {
			return 0, errors.New("Invalid trinary number")
		}
		if t != '0' {
			// cannot use math.Pow because it uses float64, so it overflows
			decimal += int64(t-'0') * pow(3, int64(len-i-1))
		}
	}
	if decimal < 0 {
		return 0, errors.New("Overflow")
	}
	return decimal, nil
}

// Using the algorithm from Donald Knuth, The Art of Computer Programming, Volume 2, Section 4.6.3
func pow(a, b int64) int64 {
	p := int64(1)
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}
