// Package prime calculates prime factors
package prime

const testVersion = 2

// Factors returns a list of prime factors for a number
func Factors(n int64) []int64 {
	factors := make([]int64, 0)
	divisor := int64(2)
	for rest := n; rest > 0; {
		divisor = factor(rest, divisor)
		if divisor == 0 {
			return factors
		}
		factors = append(factors, divisor)
		rest = rest / divisor
	}
	return factors
}

func factor(n, divisor int64) int64 {
	for i := divisor; i <= n; i++ {
		if n%i == 0 {
			return i
		}
	}
	return 0
}
