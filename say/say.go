// Package say says numbers in english
package say

import (
	"math"
	"strings"
)

var powers = []string{
	"",
	"thousand",
	"million",
	"billion",
	"trillion",
	"quadrillion",
	"quintillion",
}

var translations = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

// Say converts an integer in an english sentence
func Say(n uint64) string {
	if n == 0 {
		return "zero"
	}

	pieces := make([]string, 0)

	for j := len(powers) - 1; j >= 0; j-- {
		power := uint64(math.Pow(10, float64(j*3))) // 10^(j*3)
		group := n / power
		if group > 0 {
			pieces = append(pieces, translateHundreds(int(group)))
			if powers[j] != "" {
				pieces = append(pieces, powers[j])
			}
			n -= group * power
		}
	}

	return strings.Join(pieces, " ")
}

func translateHundreds(n int) string {
	pieces := make([]string, 0)
	if n >= 100 {
		hundreds := n / 100
		pieces = append(pieces, translateTens(hundreds)+" hundred")
		n -= hundreds * 100
	}

	if n > 0 {
		pieces = append(pieces, translateTens(n))
	}

	return strings.Join(pieces, " ")
}

func translateTens(n int) string {
	if n <= 20 || n%10 == 0 {
		return translations[n]
	}
	return translations[n/10*10] + "-" + translations[n%10]
}
