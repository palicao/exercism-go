// Package diffsquares finds the difference between the sum of the squares and the square of the sums
// of the first N natural numbers.
package diffsquares

// Difference finds the difference between the sum of the squares and the square of the sums
// of the first N natural numbers.
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}

func SumOfSquares(n int) int {
	sum := 0
	for ; n > 0; n-- {
		sum += n * n
	}
	return sum
}

func SquareOfSums(n int) int {
	sum := 0
	for ; n > 0; n-- {
		sum += n
	}
	return sum * sum
}
