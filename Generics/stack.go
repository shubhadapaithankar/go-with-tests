package stack

type Stack[T any] struct {
	values []T
}

// creates a new stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds an item to the top of the stack.
func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.values) == 0 {
		var zero T // Get zero value of type T.
		return zero, false
	}
	index := len(s.values) - 1
	element := s.values[index]
	s.values = s.values[:index]
	return element, true
}

// checks if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}
