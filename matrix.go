package main

import (
	"strconv"
	"strings"
)

type Matrix struct {
	rows    int
	columns int
	order   string
	matrix  matrix
}

type matrix [][]int

// NewMatrix create new instance of 2D matrix.
func NewMatrix(mtx matrix) *Matrix {
	rows := len(mtx)
	columns := len(mtx[0])
	m := &Matrix{
		rows:    rows,
		columns: columns,
		matrix:  mtx,
	}
	return m
}

// Rows returns number of matrix rows.
func (m *Matrix) Rows() int {
	return m.rows
}

// Columns returns number of matrix columns.
func (m *Matrix) Columns() int {
	return m.columns
}

// Show format the matrix, so it is easy for you to read.
func (m *Matrix) Show() string {
	sb := &strings.Builder{}
	for _, row := range m.matrix {
		sb.WriteString("[")
		for _, column := range row {
			sb.WriteString("\t")
			sb.WriteString(strconv.Itoa(column))
		}
		sb.WriteString("\t]")
		sb.WriteString("\n")
	}
	return sb.String()
}

// Order returns the matrix order.
// The order of a matrix is defined in terms of its number of rows and columns.
// Order of a matrix = No. of rows × No. of columns.
func (m *Matrix) Order() string {
	return "Matrix order: " + strconv.Itoa(m.rows) + " × " + strconv.Itoa(m.columns)
}

// Value returns value of matrix element at given row an column.
func (m *Matrix) Value(row, column int) int {
	return m.matrix[row-1][column-1]
}

// SetValue set value of matrix element at given row an column.
func (m *Matrix) SetValue(row, column, value int) {
	m.matrix[row-1][column-1] = value
}

// Transpose returns the matrix transpose.
// The transpose of an m x n matrix is the n x m matrix
// obtained by interchanging the rows and columns of matrix.
func (m *Matrix) Transpose() {
	transpose := m.makeMatrix(m.rows, m.columns)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.columns; j++ {
			transpose[j][i] = m.matrix[i][j]
		}
	}
	m.matrix = transpose
	transpose = nil
}

// Determinant returns the matrix determinants.
// The determinant of a matrix is a number associated with that square matrix.
// The determinant of a matrix can only be calculated for a square matrix.
// It is represented by |A|.
func (m *Matrix) Determinant() int {
	_M := m.matrix
	if m.rows == 2 && m.columns == 2 {
		// If the matrix order is 2 × 2, the determinant |A| = ad - bc
		a := _M[0][0]
		b := _M[0][1]
		c := _M[1][0]
		d := _M[1][1]
		determinant := a*d - b*c
		return determinant
	}
	// TODO: If the matrix is n × n, where n > 2, then use LU decomposition to calculate the matrix determinant.
	return 0
}

// makeMatrix will make new 2D matrix.
func (m *Matrix) makeMatrix(rows, columns int) matrix {
	mtx := make([][]int, rows)
	for i := range mtx {
		mtx[i] = make([]int, columns)
	}
	return mtx
}
