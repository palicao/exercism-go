// Package house builds the nursery song "The house that Jack built"
package house

import "strings"

const (
	subject    = "This is"
	nounPhrase = "the house that Jack built."
)

var phrases = []string{
	"the malt\nthat lay in",
	"the rat\nthat ate",
	"the cat\nthat killed",
	"the dog\nthat worried",
	"the cow with the crumpled horn\nthat tossed",
	"the maiden all forlorn\nthat milked",
	"the man all tattered and torn\nthat kissed",
	"the priest all shaven and shorn\nthat married",
	"the rooster that crowed in the morn\nthat woke",
	"the farmer sowing his corn\nthat kept",
	"the horse and the hound and the horn\nthat belonged to",
}

// Embed embeds a noun phrase as the object of relative clause with a
// transitive verb.
func Embed(relPhrase, nounPhrase string) string {
	return relPhrase + " " + nounPhrase
}

// Verse generates a verse of a song with relative clauses that have
// a recursive structure.
func Verse(subject string, relPhrases []string, nounPhrase string) string {
	if len(relPhrases) == 0 {
		return Embed(subject, nounPhrase)
	}
	return subject + " " + Embed(strings.Join(relPhrases, " "), nounPhrase)
}

// Song generates the full text of "The House That Jack Built".  Oh yes, you
// could just return a string literal, but humor us; use Verse as a subroutine.
func Song() string {
	verses := []string{}
	count := len(phrases)
	for i := 0; i <= count; i++ {
		verses = append(verses, Verse(subject, reverse(phrases[:i]), nounPhrase))
	}
	return strings.Join(verses, "\n\n")
}

// reverse returns a reversed slice
func reverse(items []string) []string {
	l := len(items)
	reversed := make([]string, l)
	for i, v := range items {
		reversed[l-i-1] = v
	}
	return reversed
}
