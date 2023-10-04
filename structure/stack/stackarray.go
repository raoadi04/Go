// Stack Array
// description: based on `geeksforgeeks` description Stack is a linear data structure which follows a particular order in which the operations are performed.
//	The order may be LIFO(Last In First Out) or FILO(First In Last Out).
// details:
// 	Stack Data Structure : https://www.geeksforgeeks.org/stack-data-structure-introduction-program/
// 	Stack (abstract data type) : https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
// author [Milad](https://github.com/miraddo)
// see stacklinkedlist.go, stacklinkedlistwithlist.go, stack_test.go

package stack

type StackArray struct {
	elements []interface{}
}

// NewStack creates and returns a new stack.
func NewStack() *StackArray {
	return &StackArray{}
}

// Push adds an element to the top of the stack.
func (s *StackArray) Push(value interface{}) {
	s.elements = append(s.elements, value)
}

// Size returns the number of elements in the stack.
func (s *StackArray) Length() int {
	return len(s.elements)
}

// Peek returns the top element of the stack without removing it.
func (s *StackArray) Peek() interface{} {
	if s.IsEmpty() {
		return nil // Stack is empty
	}
	return s.elements[len(s.elements)-1]
}

// IsEmpty returns true if the stack is empty, false otherwise.
func (s *StackArray) IsEmpty() bool {
	return len(s.elements) == 0
}

// Pop removes and returns the top element from the stack.
func (s *StackArray) Pop() interface{} {
	if s.IsEmpty() {
		return nil // Stack is empty
	}
	index := len(s.elements) - 1
	popped := s.elements[index]
	s.elements = s.elements[:index]
	return popped
}
