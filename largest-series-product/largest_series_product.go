// Package lsproduct calculates the largest product for a contiguous substring of digits of a given length
package lsproduct

import (
	"errors"
	"strconv"
	"unicode"
)

const testVersion = 3

// LargestSeriesProduct calculates the largest product for a contiguous substring of digits of a given length
func LargestSeriesProduct(digits string, span int) (product int64, err error) {
	switch {
	case span == 0:
		return 1, nil
	case span < 0:
		return 0, errors.New("span must be > 0")
	case span > len(digits):
		return 0, errors.New("span must be shorter than the lenght of digits")
	case !allNumbers(digits):
		return -1, errors.New("the digit string must contain only numbers")
	default:
		return largestProduct(digits, span), nil
	}
}

func largestProduct(digits string, span int) int64 {
	var max int64
	var product int64
	length := len(digits)
	for i := 0; i <= length-span; i++ {
		substr := digits[i : span+i]
		product = 1
		for _, d := range substr {
			multiplier, _ := strconv.Atoi(string(d))
			product *= int64(multiplier)
		}
		if product > max {
			max = product
		}
	}
	return max
}

func allNumbers(digits string) bool {
	for _, d := range digits {
		if !unicode.IsNumber(d) {
			return false
		}
	}
	return true
}
