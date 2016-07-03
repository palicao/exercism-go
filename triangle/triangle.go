// Package triangle determines the type of a triangle given the side's length
package triangle

import "math"

const testVersion = 2

type Kind int

const (
	NaT = iota
	Equ
	Iso
	Sca
)

// KindFromSides returns the type of the triangle
func KindFromSides(a, b, c float64) Kind {
	if !isTriangle(a, b, c) {
		return NaT
	}
	if a == b && a == c && b == c {
		return Equ
	}
	if a == b || a == c || b == c {
		return Iso
	}
	return Sca
}

// isTriangle checks if three sides can be a triangle
func isTriangle(a, b, c float64) bool {
	if a == 0 || b == 0 || c == 0 || math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return false
	}
	return (a+b >= c) && (a+c >= b) && (b+c >= a)
}
