package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var linkedList = NewLinkedList("Evanbill", "Antonio", "Kore", "Billy")

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
	found := linkedList.Find("Kore")
	notFound := linkedList.Find("Flo")
	assert.True(t, found)
	assert.False(t, notFound)
}

func TestNodeValue(t *testing.T) {
	val := linkedList.Value(3)
	zero := linkedList.Value(7)
	assert.NotNil(t, val)
	assert.Equal(t, zero, "")
}

func TestReverseList(t *testing.T) {
	linkedList.Print()
	linkedList.Reverse()
	linkedList.Print()
}

func TestLinkedList_Push(t *testing.T) {
	linkedList.Push("Flo")
	linkedList.Print()
}

func TestLinkedList_Append(t *testing.T) {
	linkedList.Append("Oyen")
	linkedList.Print()
}

func TestLinkedList_Insert(t *testing.T) {
	linkedList.Insert("Billy", 2)
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
