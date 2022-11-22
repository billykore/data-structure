package main

type Queue struct {
	items []int
}

func NewQueue(items ...int) *Queue {
	return &Queue{items: items}
}

func (q *Queue) Enqueue(item int) {
	q.enqueue(item)
}

func (q *Queue) Dequeue() int {
	return q.dequeue()
}

func (q *Queue) enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) dequeue() int {
	q.items = q.items[1:]
	return q.items[0]
}
