package stack_test

import (
	"testing"

	"github.com/abiiranathan/algo/stack"
)

func TestStack(t *testing.T) {
	stack := stack.New[string]()

	stack.Push("a stack")
	stack.Push("is a")
	stack.Push("lifo")
	stack.Push("data structure")

	for !stack.IsEmpty() {
		if _, empty := stack.Pop(); empty == false {
			t.Errorf("pop on stack that has elements should return true")
		}
	}

	if _, empty := stack.Pop(); empty != false {
		t.Errorf("pop on empty stack should return false")
	}

}
