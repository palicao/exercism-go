// Package summultiples sums multiplies of a given list of divisors up to a given number
package summultiples

// MultipleSummer returns a function that sums multiplies of a given list of divisors up to a given number
func MultipleSummer(divisors ...int) func(max int) int {
	return func(max int) int {
		sum := 0
		for i := 1; i < max; i++ {
			if isMultiple(i, divisors) {
				sum += i
			}
		}
		return sum
	}
}

func isMultiple(n int, divisors []int) bool {
	for _, divisor := range divisors {
		if n%divisor == 0 {
			return true
		}
	}
	return false
}
