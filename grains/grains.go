// Package grains counts grains of wheat on a chessboard
package grains

/** optimized for readability
import (
	"errors"
	"math"
)

func Square(i int) (uint64, error) {
	if i < 1 || i > 64 {
		return 0, errors.New("invalid input")
	}
	return uint64(math.Pow(2, float64(i-1))), nil
}

func Total() uint64 {
	var sum uint64
	for i := 1; i <= 64; i++ {
		cell, _ := Square(i)
		sum += cell
	}
	return sum
}
*/

/** optimized for speed */
import "errors"

func Square(i int) (uint64, error) {
	if i < 1 || i > 64 {
		return 0, errors.New("invalid input")
	}
	return 1 << uint(i-1), nil
}

func Total() uint64 {
	return 1<<64 - 1
}
