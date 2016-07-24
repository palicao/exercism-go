// Package phonenumber checks the validity of a phone number
package phonenumber

import (
	"errors"
	"unicode"
)

// Number checks and returns a normalized version of a phone number
func Number(n string) (string, error) {
	num := normalize(n)
	l := len(num)
	if l == 10 {
		return num, nil
	}
	if l == 11 && num[:1] == "1" {
		return num[1:], nil
	}
	return "", errors.New("Invalid number")
}

// AreaCode returns the area code if the number is valid
func AreaCode(n string) (string, error) {
	num, err := Number(n)
	if err == nil {
		return num[:3], nil
	}
	return "", err
}

// Format formats the number
func Format(n string) (string, error) {
	num, err := Number(n)
	if err == nil {
		return "(" + num[:3] + ") " + num[3:6] + "-" + num[6:], nil
	}
	return "", err
}

func normalize(n string) string {
	out := ""
	for _, i := range n {
		if unicode.IsNumber(i) {
			out += string(i)
		}
	}
	return out
}
