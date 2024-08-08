package adt

import "sync"

// OrderedMapError represents an error occurring from an operation performed on
// an ordered map.
type OrderedMapError string

func (e OrderedMapError) Error() string {
	return string(e)
}

var ErrDuplicateKey = OrderedMapError("duplicate key")

// OrderedMap represents an ordered collection of key-value pairs.
type OrderedMap[K comparable, V any] struct {
	order  []K
	values map[K]V
	len    int
	mu     sync.RWMutex
}

// New OrderedMap returns an initialized map.
func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		order:  []K{},
		len:    0,
		values: map[K]V{},
	}
}

// Append adds a new key-value pair to the end of the map.
func (m *OrderedMap[K, V]) Append(key K, value V) error {
	if _, has := m.values[key]; has {
		return ErrDuplicateKey
	}
	m.mu.Lock()
	defer m.mu.Unlock()

	m.values[key] = value
	m.order = append(m.order, key)
	m.len++

	return nil
}

// Delete removes the element with the specified key from the map. If there is
// no such element, then this method is a no-op.
func (m *OrderedMap[K, V]) Delete(key K) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, has := m.values[key]; !has {
		return
	}

	delete(m.values, key)

	for i, v := range m.order {
		if v == key {
			m.order = append(m.order[:i], m.order[i+1:]...)
		}
	}
	m.len--
}

// Visit invokes the given visitor function in key-value order.
func (m *OrderedMap[K, V]) Visit(visitor func(key K, value V)) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, k := range m.order {
		visitor(k, m.values[k])
	}
}

// Len returns the number of elements in the map.
func (m *OrderedMap[K, V]) Len() int {
	m.mu.RLock()
	defer m.mu.Unlock()

	return m.len
}
