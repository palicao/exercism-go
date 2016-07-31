// Package ocr recognizes numbers
package ocr

import "strings"

var numbers = map[[3]string]string{
	[3]string{
		" _ ",
		"| |",
		"|_|",
	}: "0",
	[3]string{
		"   ",
		"  |",
		"  |",
	}: "1",
	[3]string{
		" _ ",
		" _|",
		"|_ ",
	}: "2",
	[3]string{
		" _ ",
		" _|",
		" _|",
	}: "3",
	[3]string{
		"   ",
		"|_|",
		"  |",
	}: "4",
	[3]string{
		" _ ",
		"|_ ",
		" _|",
	}: "5",
	[3]string{
		" _ ",
		"|_ ",
		"|_|",
	}: "6",
	[3]string{
		" _ ",
		"  |",
		"  |",
	}: "7",
	[3]string{
		" _ ",
		"|_|",
		"|_|",
	}: "8",
	[3]string{
		" _ ",
		"|_|",
		" _|",
	}: "9",
}

func recognizeDigit(n [3]string) string {
	s, ok := numbers[n]
	if ok {
		return s
	}
	return "?"
}

func recognizeLine(row [3]string) (out string) {
	ll := len(row[0])
	for i := ll; i > 0; i -= 3 {
		var char [3]string
		for j := 0; j < 3; j++ {
			char[j] = row[j][:3]
			row[j] = row[j][3:]
		}

		out += recognizeDigit(char)
	}
	return out
}

// Recognize recognizes OCR and transforms it into strings
func Recognize(lines string) (out []string) {
	tl := lines[1:]
	rows := strings.Split(tl, "\n")
	nRows := len(rows)
	for i := 0; i < nRows; i += 4 {
		var row [3]string
		for j := 0; j < 3; j++ {
			row[j] = rows[j]
		}
		out = append(out, recognizeLine(row))
		rows = rows[4:]
	}
	return out
}
