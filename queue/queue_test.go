package queue

import (
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	q := New[int]()

	if q.Len() != 0 {
		t.Errorf("Length should be 0")
	}

	q.Enqueue(1)

	if q.Len() != 1 {
		t.Errorf("Length should be 1")
	}

	if val, _ := q.Peek(); val != 1 {
		t.Errorf("Enqueued value should be 1")
	}

	v, _ := q.Dequeue()

	if v != 1 {
		t.Errorf("Dequeued value should be 1")
	}

	_, ok := q.Peek()
	_, hasValue := q.Dequeue()
	if ok || hasValue {
		t.Errorf("Empty queue should have no values")
	}

	q.Enqueue(1)
	q.Enqueue(2)

	if q.Len() != 2 {
		t.Errorf("Length should be 2")
	}

	if value, _ := q.Peek(); value != 1 {
		t.Errorf("First value should be 1")
	}

	q.Dequeue()

	if value, _ := q.Peek(); value != 2 {
		t.Errorf("Next value should be 2")
	}

	q.Dequeue()

	if !q.Empty() {
		t.Errorf("queue should be empty")
	}
}
