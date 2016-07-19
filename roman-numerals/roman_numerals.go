// Package romannumerals coverts arabic numbers into roman
package romannumerals

import "errors"

const testVersion = 2

type arabicToRoman struct {
	arabic int
	roman  string
}

var conversion = []arabicToRoman{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// ToRomanNumeral converts an arabic number to roman
func ToRomanNumeral(arabic int) (roman string, err error) {
	if arabic < 1 || arabic > 3999 {
		return "", errors.New("numbers lower than 1 and higher than 3999 are not handled")
	}

	for rest := arabic; rest > 0; {
		part, err := convert(rest)
		if err == nil {
			roman += part.roman
			rest -= part.arabic
		}
	}

	return roman, nil
}

func convert(a int) (arabicToRoman, error) {
	for i := 0; i < len(conversion); i++ {
		if a-conversion[i].arabic >= 0 {
			return conversion[i], nil
		}
	}
	return arabicToRoman{0, ""}, errors.New("conversion not found")
}
