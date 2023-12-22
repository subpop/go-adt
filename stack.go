package adt

// Frame is an element of a stack.
type Frame[V any] struct {
	next  *Frame[V]
	Value V
}

// Stack represents an ordered collection of elements that can be accessed in a
// last-in, first-out manner. The zero value for Stack is an empty stack ready
// to use.
type Stack[V any] struct {
	head *Frame[V]
	size int
}

// NewStack returns an initalized stack.
func NewStack[V any]() *Stack[V] {
	return &Stack[V]{}
}

// Push inserts a new frame with value v on the top of stack s and returns it.
func (s *Stack[V]) Push(v V) *Frame[V] {
	f := &Frame[V]{
		Value: v,
		next:  s.head,
	}
	s.head = f
	s.size++
	return f
}

// Pop removes and returns the top-most value of stack s.
func (s *Stack[V]) Pop() *V {
	if s.head == nil {
		return nil
	}
	f := s.head
	s.head = s.head.next
	s.size--
	return &f.Value
}

// Size returns the number of frames in the stack.
func (s *Stack[V]) Size() int {
	return s.size
}
