package adt

// Element is an element of a queue.
type Element struct {
	next, prev *Element
	Value      interface{}
}

// Queue represents an ordered collection of elements that can be accessed in a
// first-in, first-out manner. The zero value for Queue is an empty queue ready
// to use.
type Queue struct {
	front, back *Element
	len         int
}

// NewQueue returns an initialized queue.
func NewQueue() *Queue {
	return &Queue{}
}

// Enqueue inserts a new element e with value v to the back of queue q and
// returns e.
func (q *Queue) Enqueue(v interface{}) *Element {
	e := &Element{
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
func (q *Queue) Dequeue() interface{} {
	e := q.front
	if q.len == 0 {
		return nil
	}
	q.front = q.front.prev
	q.len--
	return e.Value
}

// Len returns the number of elements in the queue.
func (q *Queue) Len() int {
	return q.len
}
