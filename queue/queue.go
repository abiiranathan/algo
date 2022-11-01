package queue

import "sync"

// node stores value in the queue and a reference to the next value.
type node[T any] struct {
	value T
	next  *node[T]
}

// Generic queue gaurded by a mutex.
type Queue[T any] struct {
	start, end *node[T]
	length     int

	mu sync.Mutex
}

// Create a new queue
func New[T any]() *Queue[T] {
	return &Queue[T]{start: nil, end: nil, length: 0}
}

// Take the next item off the front of the queue
// If the queue is empty, ok is false
func (queue *Queue[T]) Dequeue() (value T, ok bool) {
	queue.mu.Lock()
	defer queue.mu.Unlock()

	if queue.length == 0 {
		return value, false
	}

	n := queue.start
	if queue.length == 1 {
		queue.start = nil
		queue.end = nil
	} else {
		queue.start = queue.start.next
	}
	queue.length--

	return n.value, true
}

// Put an item on the end of a queue
func (queue *Queue[T]) Enqueue(value T) {
	queue.mu.Lock()
	defer queue.mu.Unlock()

	n := &node[T]{value: value, next: nil}

	if queue.length == 0 {
		queue.start = n
		queue.end = n
	} else {
		queue.end.next = n
		queue.end = n
	}

	queue.length++
}

// Return the number of items in the queue
func (queue *Queue[T]) Len() int {
	return queue.length
}

// Returns true is there are no items in the queue
func (queue *Queue[T]) Empty() bool {
	return queue.length == 0
}

// Return the first item in the queue without removing it
// If queue is empty, ok is false
func (queue *Queue[T]) Peek() (value T, ok bool) {
	if queue.length == 0 {
		return value, false
	}

	return queue.start.value, true
}
