package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var queue = NewQueue(1, 2, 3, 4, 5)

func TestNewQueue(t *testing.T) {
	assert.NotNil(t, queue)
	fmt.Println(queue)
}

func TestQueue_Enqueue(t *testing.T) {
	queue.Enqueue(6)
	queue.enqueue(7)
	fmt.Println(queue)
}

func TestQueue_Dequeue(t *testing.T) {
	queue.Dequeue()
	queue.Dequeue()
	fmt.Println(queue)
}
