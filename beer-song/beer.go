// Package beer sings songs about beer
package beer

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

// Verse returns one verse of the song
func Verse(i int) (v string, e error) {
	if i > 99 || i < 0 {
		return "", errors.New("The number of beer must be >= 0 and <= 99")
	}
	return fmt.Sprintf(
		"%s of beer on the wall, %s of beer.\n%s, %s of beer on the wall.\n",
		uppercaseFirst(bottle(i)),
		bottle(i),
		action(i),
		bottle(i-1),
	), nil
}

func bottle(i int) string {
	switch {
	case i > 1:
		return strconv.Itoa(i) + " bottles"
	case i == 1:
		return strconv.Itoa(i) + " bottle"
	case i == 0:
		return "no more bottles"
	default:
		return "99 bottles" // negative
	}
}

func action(i int) string {
	switch {
	case i > 1:
		return "Take one down and pass it around"
	case i == 1:
		return "Take it down and pass it around"
	default:
		return "Go to the store and buy some more"
	}
}

func uppercaseFirst(s string) string {
	return string(unicode.ToUpper(rune(s[0]))) + s[1:]
}

// Verses returns verses between upper and lower bounds
func Verses(upper, lower int) (v string, e error) {
	if lower > upper {
		return "", errors.New("The upper bound must be greater than the lower bound")
	}
	for i := upper; i >= lower; i-- {
		verse, err := Verse(i)
		if err != nil {
			return "", err
		}
		v += verse + "\n"
	}
	return v, nil
}

// Song returns the whole song
func Song() string {
	verses, _ := Verses(99, 0)
	return verses
}
