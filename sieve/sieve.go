// Package sieve implements the Sieve of Eratosthenes
package sieve

// Sieve applies the Sieve of Eratosthenes to get the first prime numbers up to a given limit
// Some optimizations can be done, this is the most basic implementation
func Sieve(limit int) []int {
	checked := make([]bool, limit+1)
	primes := []int{}
	for i := 2; i <= limit; i++ {
		if !checked[i] {
			for j := i + i; j <= limit; j += i {
				checked[j] = true
			}
			primes = append(primes, i)
		}
	}
	return primes
}
