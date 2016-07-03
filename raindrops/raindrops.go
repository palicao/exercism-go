// Package raindrops converts a number into drops!
package raindrops

import (
	"strconv"
)

const testVersion = 2

// Convert converts an integer into a string by using raindrop rules
func Convert(i int) string {
	output := ""
	if i%3 == 0 {
		output += "Pling"
	}
	if i%5 == 0 {
		output += "Plang"
	}
	if i%7 == 0 {
		output += "Plong"
	}
	if output == "" {
		output = strconv.Itoa(i)
	}
	return output
}
