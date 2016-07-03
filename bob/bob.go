// Package bob implements a lackadaisical teenager :)
package bob

import "strings"

const (
	testVersion = 2
	fine        = "Fine. Be that way!"
	whoa        = "Whoa, chill out!"
	sure        = "Sure."
	whatever    = "Whatever."
)

// Hey returns Bob's response to a sentence
func Hey(s string) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return fine
	}
	if strings.ToUpper(s) == s && strings.ToLower(s) != s {
		return whoa
	}
	if (s[len(s)-1]) == '?' {
		return sure
	}
	return whatever
}
