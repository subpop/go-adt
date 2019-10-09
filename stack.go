package adt

// Frame is an element of a stack.
type Frame struct {
	next  *Frame
	Value interface{}
}

// Stack represents an ordered collection of elements that can be accessed in a
// last-in, first-out manner. The zero value for Stack is an empty stack ready
// to use.
type Stack struct {
	head *Frame
	size int
}

// NewStack returns an initalized stack.
func NewStack() *Stack {
	return &Stack{}
}

// Push inserts a new frame with value v on the top of stack s and returns f.
func (s *Stack) Push(v interface{}) *Frame {
	f := &Frame{
		Value: v,
		next:  s.head,
	}
	s.head = f
	s.size++
	return f
}

// Pop removes and returns the top-most frame of stack s.
func (s *Stack) Pop() interface{} {
	if s.head == nil {
		return nil
	}
	f := s.head
	s.head = s.head.next
	s.size--
	return f.Value
}

// Size returns the number of frames in the stack.
func (s *Stack) Size() int {
	return s.size
}
