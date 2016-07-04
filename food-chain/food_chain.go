// Package foodchain creates beautiful songs
package foodchain

import "strings"

const (
	testVersion = 2
	nonDead     = "I don't know why she swallowed the fly. Perhaps she'll die."
	dead        = "She's dead, of course!"
)

type Animal struct {
	name        string
	comment     string
	description string
	deadly      bool
}

var animals = map[int]Animal{
	1: {name: "fly", comment: "", description: "", deadly: false},
	2: {name: "spider", comment: "It wriggled and jiggled and tickled inside her.", description: " that wriggled and jiggled and tickled inside her", deadly: false},
	3: {name: "bird", comment: "How absurd to swallow a bird!", description: "", deadly: false},
	4: {name: "cat", comment: "Imagine that, to swallow a cat!", description: "", deadly: false},
	5: {name: "dog", comment: "What a hog, to swallow a dog!", description: "", deadly: false},
	6: {name: "goat", comment: "Just opened her throat and swallowed a goat!", description: "", deadly: false},
	7: {name: "cow", comment: "I don't know how she swallowed a cow!", description: "", deadly: false},
	8: {name: "horse", comment: "", description: "", deadly: true},
}

// Verse returns a single verse of the song
func Verse(i int) string {
	out := "I know an old lady who swallowed a " + animals[i].name + ".\n"
	if animals[i].comment != "" {
		out += animals[i].comment + "\n"
	}
	if animals[i].deadly {
		out += dead
		return out
	}
	for j := i - 1; j > 0; j-- {
		out += "She swallowed the " + animals[j+1].name + " to catch the " + animals[j].name + animals[j].description + ".\n"
	}
	out += nonDead
	return out
}

// Verses returns one or more verses
func Verses(nums ...int) string {
	verses := []string{}
	for _, v := range nums {
		verses = append(verses, Verse(v))
	}
	return strings.Join(verses, "\n\n")
}

// Song returns the whole song
func Song() string {
	return Verses(1, 2, 3, 4, 5, 6, 7, 8)
}
