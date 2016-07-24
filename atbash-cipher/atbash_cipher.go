// Package atbash implements Atbash Cipher
package atbash

import "strings"

var alphaMap = alphabetMap()

// Atbash encrypts a string using the Atbash Cipher
func Atbash(s string) string {
	out := make([]rune, 0)
	count := 0
	for _, i := range s {
		encoded, ok := alphaMap[i]
		if ok {
			out = append(out, encoded)
			count++
			if count%5 == 0 {
				out = append(out, ' ')
			}
		}
	}
	return strings.TrimSpace(string(out))
}

func alphabetMap() map[rune]rune {
	runes := make(map[rune]rune)
	for uLower, e := 'a', 'z'; uLower <= 'z'; {
		runes[uLower] = e
		uLower++
		e--
	}
	for uUpper, e := 'A', 'z'; uUpper <= 'Z'; {
		runes[uUpper] = e
		uUpper++
		e--
	}
	for i := '0'; i <= '9'; i++ {
		runes[i] = i
	}
	return runes
}
