// Package igpay translates words and sentences into Pig Latin
package igpay

import (
	"regexp"
	"strings"
)

// y and x are vowels only when followed by a non-vowel
var startsWithVowel = regexp.MustCompile("^[aeiou]|^[yx][^aeiou]")

// the first letter or the first syllable if contains ch, qu, th (ex. r-une, ch-air, sch-ool, th-orn, thr-ead, ...)
var firstLetter = regexp.MustCompile(".?(ch|qu|th)[^aeiou]?|.")

// PigLatin translates a string into Pig Latin
func PigLatin(s string) string {
	words := regexp.MustCompile("\\s+").Split(s, -1)
	out := make([]string, 0)

	for _, s := range words {
		out = append(out, pigLatinWord(s))
	}

	return strings.Join(out, " ")
}

func pigLatinWord(word string) string {
	if startsWithVowel.MatchString(word) {
		return word + "ay"
	}
	l := firstLetter.FindStringIndex(word)
	return word[l[1]:] + word[l[0]:l[1]] + "ay"
}
