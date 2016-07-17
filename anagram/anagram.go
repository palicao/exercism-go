// Package anagram detects anagrams
package anagram

import "strings"

// Detect lists all anagrams for a string in a list of candidates
func Detect(s string, candidates []string) (anagrams []string) {
	s = strings.ToLower(s)
	for _, candidate := range candidates {
		candidate = strings.ToLower(candidate)
		if s == candidate || len(s) != len(candidate) {
			continue
		}
		if isAnagram(s, candidate) {
			anagrams = append(anagrams, candidate)
		}
	}
	return anagrams
}

func isAnagram(s, anagram string) bool {
	for _, l := range s {
		index := strings.IndexRune(anagram, l)
		if index == -1 {
			return false
		}
		anagram = anagram[:index] + anagram[index+1:]
	}
	return anagram == ""
}
