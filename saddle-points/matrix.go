// Package matrix builds matrices from strings
package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

// New creates a new matrix from a string
func New(s string) (*Matrix, error) {
	m := Matrix{}
	prevCols := -1
	rows := strings.Split(s, "\n")
	for _, row := range rows {
		cells := strings.Split(strings.TrimSpace(row), " ")
		l := len(cells)
		if prevCols == -1 {
			prevCols = l
		}
		if l == 0 || prevCols != l {
			return &Matrix{}, errors.New("Uneven number of columns")
		}
		intCells := make([]int, l)
		for j, cell := range cells {
			c, err := strconv.Atoi(cell)
			if err != nil {
				return &Matrix{}, errors.New("A cell contains a non-numeric value")
			}
			intCells[j] = c
		}
		m = append(m, intCells)
	}
	return &m, nil
}

// Rows returns the matrix
func (m Matrix) Rows() [][]int {
	rows := make([][]int, len(m))
	for r, row := range m {
		rows[r] = make([]int, len(row))
		for c, cell := range row {
			rows[r][c] = cell
		}
	}
	return rows
}

// Cols returns the matrix rotated by 90°
func (m Matrix) Cols() [][]int {
	cols := make([][]int, len(m[0]))
	for i := 0; i < len(m[0]); i++ {
		cols[i] = make([]int, len(m))
	}
	for r, row := range m {
		for c, cell := range row {
			cols[c][r] = cell
		}
	}
	return cols
}

// Set sets an element in the matrix
func (m Matrix) Set(r, c, val int) bool {
	if r >= len(m) || r < 0 || c >= len(m[0]) || c < 0 {
		return false
	}
	m[r][c] = val
	return true
}
