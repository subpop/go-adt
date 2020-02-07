package adt

import "container/heap"

// PriorityQueue represents an ordered collection of elements that can be
// accessed in a priority-first manner. Elements with a higher priority are
// retrieved before elements with a lower priority. Priority is represented in
// a lower-value order; the closer a priority is to zero, the higher priority
// it receives.
type PriorityQueue struct {
	elements *pqueue
}

// NewPriorityQueue returns a priority queue of size n.
func NewPriorityQueue(n int) *PriorityQueue {
	elements := make(pqueue, n)
	pq := &PriorityQueue{
		elements: &elements,
	}
	heap.Init(pq.elements)
	return pq
}

// Push inserts a new element e with value v and priority p to the queue.
func (pq *PriorityQueue) Push(v interface{}, p float64) {
	e := &element{
		value:    v,
		priority: p,
	}
	heap.Push(pq.elements, e)
}

// Pop removes the element with the highest priority and returns its value.
func (pq *PriorityQueue) Pop() interface{} {
	if len(*pq.elements) == 0 {
		return nil
	}
	elem := heap.Pop(pq.elements).(*element)
	return elem.value
}

// Len returns the number of elements in the queue.
func (pq PriorityQueue) Len() int {
	return len(*pq.elements)
}

type element struct {
	value    interface{}
	priority float64
}

type pqueue []*element

func (pq pqueue) Len() int {
	return len(pq)
}

func (pq pqueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq pqueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *pqueue) Push(x interface{}) {
	elem := x.(*element)
	*pq = append(*pq, elem)
}

func (pq *pqueue) Pop() interface{} {
	old := *pq
	n := len(old)
	elem := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return elem
}
