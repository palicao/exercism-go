// Package hexadecimal converts hex to decimal
package hexadecimal

import (
	"errors"
	"strings"
)

// ParseHex converts hex to decimal
func ParseHex(hex string) (dec int64, err error) {
	l := len(hex)
	if l < 1 {
		return 0, errors.New("syntax")
	}
	hex = strings.ToLower(hex)
	for i, h := range hex {
		if (h < '0' || h > '9') && (h < 'a' || h > 'f') {
			return 0, errors.New("syntax")
		}

		exp := int64(l - i - 1)
		if exp > 14 && h != '0' {
			return 0, errors.New("range")
		}

		if h >= '1' && h <= '9' {
			dec += int64(h-'0') * pow(16, exp)
		} else if h >= 'a' && h <= 'f' {
			dec += int64(h-'a'+10) * pow(16, exp)
		}
	}
	return dec, nil
}

// HandleErrors returns a list of errors from a list of hexadecimal numbers
func HandleErrors(hexList []string) (errorStrings []string) {
	for _, hex := range hexList {
		errorStrings = append(errorStrings, handleError(hex))
	}
	return errorStrings
}

func handleError(hex string) string {
	_, err := ParseHex(hex)
	if err == nil {
		return "none"
	}
	return err.Error()
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
