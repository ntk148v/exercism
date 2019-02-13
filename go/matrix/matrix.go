package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Matrix stores an array of numbers in rows and columns
type Matrix struct {
	rows, cols int
	value      []int
}

func New(s string) (*Matrix, error) {
	rows := strings.Split(s, "\n")
	matrix := Matrix{rows: len(rows), cols: -1}
	for _, row := range rows {
		cols := strings.Split(strings.TrimSpace(row), " ")
		if matrix.cols == -1 {
			matrix.cols = len(cols)
		} else if matrix.cols != len(cols) {
			return nil, fmt.Errorf("Rows need to be the same length")
		}
		for _, c := range cols {
			num, err := strconv.Atoi(c)
			if err != nil {
				return nil, err
			}
			matrix.value = append(matrix.value, num)
		}
	}
	return &matrix, nil
}

// Set sets the value of the matrix at point row, col
func (m *Matrix) Set(row, col, val int) (ok bool) {
	if row < 0 || row >= m.rows || col < 0 || m.cols <= col {
		return false
	}
	m.value[row*m.cols+col] = val
	return true
}

// Rows gets the matrix represent in rows
func (m *Matrix) Rows() [][]int {
	rows := make([][]int, m.rows)
	for r := 0; r < m.rows; r++ {
		rows[r] = make([]int, m.cols)
		for c := 0; c < m.cols; c++ {
			rows[r][c] = m.value[r*m.cols+c]
		}
	}
	return rows
}

// Cols gets the matrix presented in columns
func (m *Matrix) Cols() [][]int {
	cols := make([][]int, m.cols)
	for c := 0; c < m.cols; c++ {
		cols[c] = make([]int, m.rows)
		for r := 0; r < m.rows; r++ {
			cols[c][r] = m.value[r*m.cols+c]
		}
	}
	return cols
}
