package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var linkedList = NewLinkedList(1, 2, 3, 4)

func TestNewLinkedList(t *testing.T) {
	assert.NotNil(t, linkedList)
}

func TestPrintLinkedList(t *testing.T) {
	linkedList.Print()
}

func TestLinkedList_Len(t *testing.T) {
	length := linkedList.Len()
	assert.Equal(t, 4, length)
}

func TestFindLinkedList(t *testing.T) {
	found := linkedList.Find(3)
	notFound := linkedList.Find(7)
	assert.True(t, found)
	assert.False(t, notFound)
}

func TestNodeValue(t *testing.T) {
	val := linkedList.Value(3)
	zero := linkedList.Value(7)
	assert.NotNil(t, val)
	assert.Equal(t, zero, 0)
}

func TestReverseList(t *testing.T) {
	linkedList.Print()
	linkedList.Reverse()
	linkedList.Print()
}

func TestLinkedList_Push(t *testing.T) {
	linkedList.Push(7)
	linkedList.Print()
}

func TestLinkedList_Append(t *testing.T) {
	linkedList.Append(7)
	linkedList.Print()
}

func TestLinkedList_Insert(t *testing.T) {
	linkedList.Insert(7, 2)
	linkedList.Print()
}

func TestLinkedList_Pull(t *testing.T) {
	linkedList.Pull()
	linkedList.Print()
}

func TestLinkedList_Delete(t *testing.T) {
	linkedList.Delete(2)
	linkedList.Print()
}

func TestLinkedList_Pop(t *testing.T) {
	linkedList.Pop()
	linkedList.Print()
}
