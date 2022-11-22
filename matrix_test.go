package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var mtx = NewMatrix(
	[][]int{
		{1, 2},
		{3, 4},
	},
)

func TestNewMatrix(t *testing.T) {
	assert.NotNil(t, mtx)
	fmt.Println(mtx.Show())
}

func TestMatrix_Rows(t *testing.T) {
	assert.Equal(t, 2, mtx.Rows())
}

func TestMatrix_Column(t *testing.T) {
	assert.Equal(t, 2, mtx.Columns())
}

func TestMatrix_Order(t *testing.T) {
	assert.Equal(t, "Matrix order: 2 × 2", mtx.Order())
}

func TestMatrix_Value(t *testing.T) {
	assert.Equal(t, 3, mtx.Value(2, 1))
}

func TestMatrix_SetValue(t *testing.T) {
	mtx.SetValue(2, 2, 100)
	fmt.Println(mtx.Show())
}

func TestMatrix_Transpose(t *testing.T) {
	mtx.Transpose()
	fmt.Println(mtx.Show())
}

func TestMatrix_Determinant(t *testing.T) {
	determinant := mtx.Determinant()
	assert.Equal(t, -2, determinant)
}
