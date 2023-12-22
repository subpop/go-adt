package adt

// Element is an element of a queue.
type Element[V any] struct {
	next, prev *Element[V]
	Value      V
}

// Queue represents an ordered collection of elements that can be accessed in a
// first-in, first-out manner. The zero value for Queue is an empty queue ready
// to use.
type Queue[V any] struct {
	front, back *Element[V]
	len         int
}

// NewQueue returns an initialized queue.
func NewQueue[V any]() *Queue[V] {
	return &Queue[V]{}
}

// Enqueue inserts a new element e with value v to the back of queue q and
// returns e.
func (q *Queue[V]) Enqueue(v V) *Element[V] {
	e := &Element[V]{
		Value: v,
		next:  q.back,
	}
	if q.len == 0 {
		q.front = e
	} else {
		q.back.prev = e
	}
	q.back = e
	q.len++
	return e
}

// Dequeue removes and returns the front-most element of queue q.
func (q *Queue[V]) Dequeue() *V {
	e := q.front
	if q.len == 0 {
		return nil
	}
	q.front = q.front.prev
	q.len--
	return &e.Value
}

// Len returns the number of elements in the queue.
func (q *Queue[V]) Len() int {
	return q.len
}
