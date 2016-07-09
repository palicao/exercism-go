// Package etl transform crufty data into shiny data
package etl

import "strings"

// Transform transforms the legacy data format into the new format
func Transform(legacy map[int][]string) map[string]int {
	newData := map[string]int{}
	for value, letters := range legacy {
		for _, letter := range letters {
			newData[strings.ToLower(letter)] = value
		}
	}
	return newData
}
