// Package dna computes how many times each nucleotide occurs in a DNA string
package dna

import "errors"

type Histogram map[byte]int

type strand struct {
	dnaString string
	histogram Histogram
}

type dna interface {
	Counts() Histogram
	Count(byte) int
}

// DNA converts a string into a strand
func DNA(d string) strand {
	h := map[byte]int{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}
	for _, n := range d {
		byteN := runeToByte(n)
		h[byteN]++
	}
	return strand{dnaString: d, histogram: h}
}

func runeToByte(r rune) byte {
	return byte(string(r)[0])
}

// Counts returns the Histogram of nucleotides in a strand
func (s *strand) Counts() Histogram {
	return s.histogram
}

// Count returns the number of nucleotides of type n found in a strand
func (s *strand) Count(n byte) (int, error) {
	num, ok := s.histogram[n]
	if !ok {
		return 0, errors.New("Error parsing strand")
	}
	return num, nil
}
