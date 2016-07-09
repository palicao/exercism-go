// Package pythagorean works on Pythagorean triplets
package pythagorean

import "math"

type Triplet [3]int

// Range returns all the triplets containing numbers between min and max
func Range(min, max int) []Triplet {
	triplets := []Triplet{}
	for _, t := range allTriplets(min, max) {
		if isPythagoreanTriplet(t) {
			triplets = append(triplets, t)
		}
	}
	return triplets
}

// Sum returns all the triplets containing numbers whose sum is p
func Sum(p int) []Triplet {
	triplets := []Triplet{}
	for _, t := range tripletsWhoseSumIs(p) {
		if isPythagoreanTriplet(t) {
			triplets = append(triplets, t)
		}
	}
	return triplets
}

func isPythagoreanTriplet(t Triplet) bool {
	return t[0]*t[0]+t[1]*t[1] == t[2]*t[2]
}

func allTriplets(min, max int) []Triplet {
	triplets := []Triplet{}
	for a := min; a <= max; a++ {
		for c := max; c >= a; c-- {
			b := int(math.Sqrt(float64(c*c - a*a)))
			if b >= a && b <= c {
				triplets = append(triplets, Triplet{a, b, c})
			}
		}
	}
	return triplets
}

func tripletsWhoseSumIs(p int) []Triplet {
	triplets := []Triplet{}
	for a := 1; a <= p/3; a++ {
		for c := p / 2; c >= a; c-- {
			b := p - c - a
			if b >= a && b <= c {
				triplets = append(triplets, Triplet{a, b, c})
			}
		}
	}
	return triplets
}
