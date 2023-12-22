package adt

import "container/heap"

// PriorityQueue represents an ordered collection of elements that can be
// accessed in a priority-first manner. Elements with a higher priority are
// retrieved before elements with a lower priority. Priority is represented in
// a lower-value order; the closer a priority is to zero, the higher priority
// it receives.
type PriorityQueue[V comparable] struct {
	elements *pqueue[V]
	indexes  map[V]int
}

// NewPriorityQueue returns a priority queue of size n.
func NewPriorityQueue[V comparable](n int) *PriorityQueue[V] {
	elements := make(pqueue[V], n)
	pq := &PriorityQueue[V]{
		elements: &elements,
		indexes:  make(map[V]int),
	}
	heap.Init(pq.elements)
	return pq
}

// Push inserts a new element e with value v and priority p to the queue.
func (pq *PriorityQueue[V]) Push(v V, p float64) {
	e := &element[V]{
		value:    v,
		priority: p,
	}
	pq.indexes[v] = len(*pq.elements)
	heap.Push(pq.elements, e)
}

// Pop removes the element with the highest priority and returns its value.
func (pq *PriorityQueue[V]) Pop() *V {
	if len(*pq.elements) == 0 {
		return nil
	}
	elem := heap.Pop(pq.elements).(*element[V])
	delete(pq.indexes, elem.value)
	return &elem.value
}

// Prioritize adjusts the priority of element v to priority p and adjusts the
// queue order.
func (pq *PriorityQueue[V]) Prioritize(v V, p float64) {
	i := pq.indexes[v]
	e := pq.elements.Elem(i)
	e.priority = p
	heap.Fix(pq.elements, i)
}

// Len returns the number of elements in the queue.
func (pq PriorityQueue[V]) Len() int {
	return len(*pq.elements)
}

type element[V any] struct {
	value    V
	priority float64
}

type pqueue[V any] []*element[V]

func (pq pqueue[V]) Len() int {
	return len(pq)
}

func (pq pqueue[V]) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq pqueue[V]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *pqueue[V]) Push(x interface{}) {
	elem := x.(*element[V])
	*pq = append(*pq, elem)
}

func (pq *pqueue[V]) Pop() interface{} {
	old := *pq
	n := len(old)
	elem := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return elem
}

func (pq pqueue[V]) Elem(i int) *element[V] {
	return pq[i]
}
