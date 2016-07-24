// Package prime tells you which is the nth prime
package prime

import "math"

// Nth returns the nth prime number
func Nth(n int) (prime int, ok bool) {
	if n < 1 {
		return 0, false
	}
	i := 2
	for foundPrimes := 0; foundPrimes < n; i++ {
		if isPrime(i) {
			foundPrimes++
		}
		if foundPrimes == n {
			return i, true
		}
	}
	return 0, false
}

func isPrime(n int) bool {
	if n == 2 {
		return true
	}
	// pre-filter all multiples of 2
	if n%2 == 0 {
		return false
	}
	// we can stop counting at the square root (n = sqrt(n) * sqrt(n))
	sq := int(math.Sqrt(float64(n)))
	// since multiples of 2 are pre-filtered we can start from 3 and exclude even numbers
	for i := 3; i <= sq; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
