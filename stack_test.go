package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var stack = NewStack(1, 2, 3)

func TestNewStack(t *testing.T) {
	assert.NotNil(t, stack)
	fmt.Println(stack)
}

func TestStack_Push(t *testing.T) {
	stack.Push(7)
	stack.Push(9)
	stack.Push(11)
	fmt.Println(stack)
}

func TestStack_Pop(t *testing.T) {
	stack.Pop()
	fmt.Println(stack)
}
