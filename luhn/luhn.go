// Package luhn calculates Luhn checksum
package luhn

import (
	"strconv"
	"unicode"
)

// Valid checks if a string contains a valid according to the Luhn formula
func Valid(code string) bool {
	code = normalize(code)
	l := len(code)
	if l < 2 {
		return false
	}
	return AddCheck(code[:l-1]) == code
}

// AddCheck calculates the checksum to append to a string to make it valid
func AddCheck(code string) string {
	check := normalize(code)
	l := len(check)
	for i := l; i > 0; i = i - 2 {
		num, _ := strconv.Atoi(check[i-1 : i])
		num *= 2
		if num > 9 {
			num -= 9
		}
		check = check[:i-1] + strconv.Itoa(num) + check[i:]
	}
	sum := 0
	for _, i := range check {
		add, _ := strconv.Atoi(string(i))
		sum += add
	}

	nextDec := (sum/10 + 1) * 10
	checksum := nextDec - sum
	if checksum == 10 {
		checksum = 0
	}
	return code + strconv.Itoa(checksum)
}

func normalize(code string) string {
	var normalized string
	for _, i := range code {
		if unicode.IsNumber(i) {
			normalized += string(i)
		}
	}
	return normalized
}
