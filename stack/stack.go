// stack implementation using a slice.
package stack

// Stack data structure
type Stack[T any] struct {
	data []T
}

// Initialize a new Stack.
func New[T any]() *Stack[T] {
	return &Stack[T]{
		data: []T{},
	}
}

// IsEmpty: check if stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Push a new value onto the stack
func (s *Stack[T]) Push(element T) {
	(*s).data = append((*s).data, element)
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack[T]) Pop() (elem T, empty bool) {
	if s.IsEmpty() {
		return elem, false
	} else {
		index := len((*s).data) - 1   // Get the index of the top most element.
		element := (*s).data[index]   // Index into the slice and obtain the element.
		(*s).data = (*s).data[:index] // Remove it from the stack by slicing it off.
		return element, true
	}
}
