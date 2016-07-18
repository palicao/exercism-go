// Package strand finds the RNA complement to a DNA strand
package strand

const testVersion = 3

var transcript = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA returns the RNA complement to a DNA strand
func ToRNA(dna string) (rna string) {
	for _, n := range dna {
		rna += string(transcript[n])
	}
	return rna
}
