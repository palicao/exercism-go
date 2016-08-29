// Package cipher implements 3 different simple cipher algorithms
package cipher

import (
	"regexp"
	"strings"
)

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type Caesar struct {
}

type Shift struct {
	shift int
}

type Vigenere struct {
	key string
}

// NewCaesar instantiates a new Caesar cipher
func NewCaesar() *Caesar {
	return &Caesar{}
}

// NewCaesar instantiates a new Shift cipher
func NewShift(shift int) *Shift {
	if shift == 0 || shift > 25 || shift < -25 {
		return nil
	}
	return &Shift{shift}
}

// NewCaesar instantiates a new Vigenere cipher
func NewVigenere(key string) *Vigenere {
	if !checkVigenereKey(key) {
		return nil
	}
	return &Vigenere{key}
}

func checkVigenereKey(key string) bool {
	totalDiff := 0
	for _, l := range key {
		if l < 'a' || l > 'z' {
			return false
		}
		// avoid an empty key or containing only one or more "a"
		totalDiff += int(l - 'a')
	}
	return totalDiff > 0
}

func prepareMessage(m string) string {
	m = strings.ToLower(m)
	re := regexp.MustCompile("[^a-z]")
	return re.ReplaceAllString(m, "")
}

func shift(r rune, diff int) rune {
	n := r + rune(diff)
	if n > 'z' {
		n -= 'z' - 'a' + 1
	}
	if n < 'a' {
		n += 'z' - 'a' + 1
	}
	return n
}

func shiftString(s string, diff int) string {
	shifted := ""
	for _, l := range s {
		shifted += string(shift(l, diff))
	}
	return shifted
}

func vigenereShift(s string, key string, encode bool) string {
	shifted := ""
	j := 0
	for _, l := range s {
		letterShift := int(key[j] - 'a')
		if !encode {
			letterShift = -letterShift
		}
		shifted += string(shift(l, letterShift))
		j++
		if j >= len(key) {
			j = 0
		}
	}
	return shifted
}

// Encode for Caesar
func (c Caesar) Encode(s string) string {
	return shiftString(prepareMessage(s), 3)
}

// Decode for Caesar
func (c Caesar) Decode(s string) string {
	return shiftString(s, -3)
}

// Encode for Shift
func (c Shift) Encode(s string) string {
	return shiftString(prepareMessage(s), c.shift)
}

// Decode for Shift
func (c Shift) Decode(s string) string {
	return shiftString(s, -c.shift)
}

// Encode for Vigenere
func (c Vigenere) Encode(s string) string {
	return vigenereShift(prepareMessage(s), c.key, true)
}

// Decode for Vigenere
func (c Vigenere) Decode(s string) string {
	return vigenereShift(prepareMessage(s), c.key, false)
}
