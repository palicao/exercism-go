// Package wordcount counts instances of each word in a sentence
package wordcount

import (
	"strings"
	"unicode"
)

const testVersion = 2

type Frequency map[string]int

// WordCount counts instances of each word in a sentence
func WordCount(phrase string) Frequency {
	f := make(Frequency)
	for _, v := range strings.FieldsFunc(phrase, func(r rune) bool {
		return !(unicode.IsLetter(r) || unicode.IsNumber(r))
	}) {
		f[strings.ToLower(v)]++
	}
	return f
}
