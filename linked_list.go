package main

import (
	"fmt"
)

type Node[T comparable] struct {
	Value T
	next  *Node[T]
}

type LinkedList[T comparable] struct {
	Head *Node[T]
	len  int
}

func (ll *LinkedList[T]) Print() {
	current := ll.Head
	for current != nil {
		if current.next != nil {
			fmt.Print(current.Value, " -> ")
		} else {
			fmt.Println(current.Value)
		}
		current = current.next
	}
}

func (ll *LinkedList[T]) Len() int {
	current := ll.Head
	for current != nil {
		ll.len++
		current = current.next
	}
	return ll.len
}

func (ll *LinkedList[T]) Find(target T) bool {
	current := ll.Head
	for current != nil {
		if current.Value == target {
			return true
		}
		current = current.next
	}
	return false
}

func (ll *LinkedList[T]) Value(index int) T {
	current := ll.Head
	for current != nil {
		if ll.len == index {
			return current.Value
		}
		current = current.next
	}
	return *new(T)
}

func (ll *LinkedList[T]) Reverse() {
	current := ll.Head
	ll.Head = new(Node[T])
	for current != nil {
		next := current.next
		current.next = ll.Head
		ll.Head = current
		current = next
	}
}

func (ll *LinkedList[T]) Push(data T) {
	head := &Node[T]{Value: data}
	head.next = ll.Head
	ll.Head = head
}

func (ll *LinkedList[T]) Append(data T) {
	newNode := &Node[T]{Value: data}
	tail := ll.Head
	for tail.next != nil {
		tail = tail.next
	}
	tail.next = newNode
}

func (ll *LinkedList[T]) Insert(data T, index int) {
	if index > ll.Len() {
		fmt.Println("index out of list range")
		return
	}
	if index == 0 {
		ll.Push(data)
		return
	}
	current := ll.Head
	count := 1
	for current != nil {
		if count == index {
			newNode := &Node[T]{Value: data}
			newNode.next = current.next
			current.next = newNode
			return
		}
		current = current.next
		count++
	}
}

func (ll *LinkedList[T]) Pull() {
	ll.Head = ll.Head.next
}

func (ll *LinkedList[T]) Pop() {
	tail := ll.Head
	for tail.next != nil {
		if tail.next.next == nil {
			tail.next = nil
			return
		}
		tail = tail.next
	}
}

func (ll *LinkedList[T]) Delete(index int) {
	if index > ll.Len() {
		fmt.Println("index out of list range")
		return
	}
	if index == 0 {
		ll.Pull()
	}
	current := ll.Head
	count := 1
	for current != nil {
		if count == index {
			current.next = current.next.next
			return
		}
		current = current.next
		count++
	}
}
