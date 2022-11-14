package main

import (
	"fmt"
)

type node[T comparable] struct {
	value T
	next  *node[T]
}

type LinkedList[T comparable] struct {
	Head   *node[T]
	length int
}

func NewLinkedList[T comparable](head T, values ...T) *LinkedList[T] {
	linkedList := &LinkedList[T]{
		Head: &node[T]{value: head},
	}
	for _, value := range values {
		linkedList.append(value)
	}
	return linkedList
}

func (ll *LinkedList[T]) Print() {
	current := ll.Head
	for current != nil {
		if current.next != nil {
			fmt.Print(current.value, " -> ")
		} else {
			fmt.Println(current.value)
		}
		current = current.next
	}
}

func (ll *LinkedList[T]) Len() int {
	return ll.len()
}

func (ll *LinkedList[T]) Find(target T) bool {
	current := ll.Head
	for current != nil {
		if current.value == target {
			return true
		}
		current = current.next
	}
	return false
}

func (ll *LinkedList[T]) Value(index int) T {
	current := ll.Head
	for current != nil {
		if ll.length == index {
			return current.value
		}
		current = current.next
	}
	return *new(T)
}

func (ll *LinkedList[T]) Reverse() {
	current := ll.Head
	ll.Head = new(node[T])
	for current != nil {
		next := current.next
		current.next = ll.Head
		ll.Head = current
		current = next
	}
}

func (ll *LinkedList[T]) Push(data T) {
	ll.push(data)
}

func (ll *LinkedList[T]) Append(data T) {
	ll.append(data)
}

func (ll *LinkedList[T]) Insert(data T, index int) {
	if index > ll.len() {
		fmt.Println("index out of list range")
		return
	}
	if index == 0 {
		ll.push(data)
		return
	}
	current := ll.Head
	count := 1
	for current != nil {
		if count == index {
			newNode := &node[T]{value: data}
			newNode.next = current.next
			current.next = newNode
			return
		}
		current = current.next
		count++
	}
}

func (ll *LinkedList[T]) Pull() {
	ll.pull()
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
	if index > ll.len() {
		fmt.Println("index out of list range")
		return
	}
	if index == 0 {
		ll.pull()
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

func (ll *LinkedList[T]) len() int {
	current := ll.Head
	for current != nil {
		ll.length++
		current = current.next
	}
	return ll.length
}

func (ll *LinkedList[T]) push(data T) {
	head := &node[T]{value: data}
	head.next = ll.Head
	ll.Head = head
}

func (ll *LinkedList[T]) append(data T) {
	newNode := &node[T]{value: data}
	tail := ll.Head
	for tail.next != nil {
		tail = tail.next
	}
	tail.next = newNode
}

func (ll *LinkedList[T]) pull() {
	ll.Head = ll.Head.next
}
