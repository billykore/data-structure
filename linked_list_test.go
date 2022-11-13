package main

import (
	"fmt"
	"testing"
)

func TestPrintLinkedList(t *testing.T) {
	one := &Node[int]{Value: 1}
	two := &Node[int]{Value: 2}
	three := &Node[int]{Value: 3}
	four := &Node[int]{Value: 4}
	five := &Node[int]{Value: 5}

	one.next = two
	two.next = three
	three.next = four
	four.next = five

	linkedList := &LinkedList[int]{Head: one}
	linkedList.Print()
}

func TestFindLinkedList(t *testing.T) {
	one := &Node[int]{Value: 1}
	two := &Node[int]{Value: 2}
	three := &Node[int]{Value: 3}
	four := &Node[int]{Value: 4}
	five := &Node[int]{Value: 5}

	one.next = two
	two.next = three
	three.next = four
	four.next = five

	linkedList := &LinkedList[int]{Head: one}
	found1 := linkedList.Find(3)
	found2 := linkedList.Find(7)

	fmt.Println(found1)
	fmt.Println(found2)
}

func TestNodeValue(t *testing.T) {
	a := &Node[rune]{Value: 'A'}
	b := &Node[rune]{Value: 'B'}
	c := &Node[rune]{Value: 'C'}
	d := &Node[rune]{Value: 'D'}

	a.next = b
	b.next = c
	c.next = d

	linkedList := &LinkedList[rune]{Head: a}
	value1 := linkedList.Value(3)
	value2 := linkedList.Value(7)

	fmt.Println(string(value1))
	fmt.Println(value2)
}

func TestReverseList(t *testing.T) {
	one := &Node[int]{Value: 1}
	two := &Node[int]{Value: 2}
	three := &Node[int]{Value: 3}
	four := &Node[int]{Value: 4}
	five := &Node[int]{Value: 5}

	one.next = two
	two.next = three
	three.next = four
	four.next = five

	linkedList := &LinkedList[int]{Head: one}
	linkedList.Reverse()
	linkedList.Print()
}

func TestLinkedList_Add(t *testing.T) {
	head := &Node[int]{Value: 1}
	linkedList := &LinkedList[int]{Head: head}

	linkedList.Push(2)
	linkedList.Push(3)
	linkedList.Push(4)
	linkedList.Push(5)
	linkedList.Print()
}

func TestLinkedList_Append(t *testing.T) {
	head := &Node[int]{Value: 1}
	linkedList := &LinkedList[int]{Head: head}

	linkedList.Append(3)
	linkedList.Append(2)
	linkedList.Append(4)
	linkedList.Print()
	fmt.Println(linkedList.Len())
}

func TestLinkedList_Insert(t *testing.T) {
	head := &Node[int]{Value: 1}
	linkedList := &LinkedList[int]{Head: head}

	linkedList.Append(3)
	linkedList.Append(2)
	linkedList.Append(4)
	linkedList.Print()

	linkedList.Insert(7, 2)
	linkedList.Print()
}

func TestLinkedList_Pull(t *testing.T) {
	head := &Node[int]{Value: 1}
	linkedList := &LinkedList[int]{Head: head}

	linkedList.Append(3)
	linkedList.Append(2)
	linkedList.Append(4)
	linkedList.Print()

	linkedList.Pull()
	linkedList.Print()
}

func TestLinkedList_Delete(t *testing.T) {
	head := &Node[int]{Value: 1}
	linkedList := &LinkedList[int]{Head: head}

	linkedList.Append(3)
	linkedList.Append(2)
	linkedList.Append(4)
	linkedList.Print()

	linkedList.Delete(2)
	linkedList.Print()
}

func TestLinkedList_Pop(t *testing.T) {
	head := &Node[int]{Value: 1}
	linkedList := &LinkedList[int]{Head: head}

	linkedList.Append(3)
	linkedList.Append(2)
	linkedList.Append(4)
	linkedList.Print()

	linkedList.Pop()
	linkedList.Print()
}
