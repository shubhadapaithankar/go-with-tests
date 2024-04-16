package stack

import (
	"testing"
)

// TestStackOperations tests various operations on the Stack type.
func TestStackOperations(t *testing.T) {
	t.Run("Integer Stack", func(t *testing.T) {
		intStack := NewStack[int]()
		if !intStack.IsEmpty() {
			t.Error("New stack should be empty")
		}

		intStack.Push(10)
		intStack.Push(20)

		if value, ok := intStack.Pop(); !ok || value != 20 {
			t.Errorf("Pop should return last pushed item 20, got %d", value)
		}

		if value, ok := intStack.Pop(); !ok || value != 10 {
			t.Errorf("Pop should return last pushed item 10, got %d", value)
		}

		if _, ok := intStack.Pop(); ok {
			t.Error("Pop should fail on empty stack")
		}
	})

	t.Run("String Stack", func(t *testing.T) {
		stringStack := NewStack[string]()
		stringStack.Push("hello")
		stringStack.Push("world")

		if value, ok := stringStack.Pop(); !ok || value != "world" {
			t.Errorf("Pop should return last pushed item 'world', got %s", value)
		}

		if value, ok := stringStack.Pop(); !ok || value != "hello" {
			t.Errorf("Pop should return last pushed item 'hello', got %s", value)
		}

		if _, ok := stringStack.Pop(); ok {
			t.Error("Pop should fail on empty stack")
		}
	})
}
