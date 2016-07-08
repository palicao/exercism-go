// Package queenattack calculates if two queens can attack each others
package queenattack

import "errors"

// CanQueenAttack checks if the white and the black queen can attack each other
func CanQueenAttack(w, b string) (bool, error) {
	if !hasCorrectNotation(w) || !hasCorrectNotation(b) || w == b {
		return false, errors.New("invalid notation")
	}
	colDist := int(w[0]) - int(b[0])
	rowDist := int(w[1]) - int(b[1])
	if rowDist == 0 || colDist == 0 || rowDist == colDist || rowDist == -colDist {
		return true, nil
	}
	return false, nil
}

// hasCorrectNotation checks the algebraic notation
func hasCorrectNotation(position string) bool {
	if len(position) != 2 {
		return false
	}
	return position[0] >= 'a' && position[1] <= 'h' && position[1] >= '1' && position[1] <= '8'
}
