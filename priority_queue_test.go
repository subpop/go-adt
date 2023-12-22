package adt

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPriorityQueueString(t *testing.T) {
	tests := []struct {
		input      map[string]float64
		want       []string
		prioritize map[string]float64
	}{
		{
			input: map[string]float64{
				"a": 3,
				"b": 2,
				"c": 5,
				"d": 0,
			},
			want: []string{"d", "b", "a", "c"},
		},
		{
			input: map[string]float64{},
			want:  []string{},
		},
		{
			input: map[string]float64{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			want: []string{"c", "a", "b"},
			prioritize: map[string]float64{
				"c": 0,
			},
		},
	}

	for _, test := range tests {
		q := NewPriorityQueue[string](0)
		for v, p := range test.input {
			q.Push(v, p)
		}

		if q.Len() != len(test.input) {
			t.Fatalf("%v != %v", q.Len(), len(test.input))
		}

		for v, p := range test.prioritize {
			q.Prioritize(v, p)
		}

		for _, want := range test.want {
			got := q.Pop()
			if !cmp.Equal(*got, want) {
				t.Errorf("%v != %v", *got, want)
			}
		}
	}
}

func TestPriorityQueueInt(t *testing.T) {
	tests := []struct {
		input      map[int]float64
		want       []int
		prioritize map[int]float64
	}{
		{
			input: map[int]float64{
				1: 3,
				2: 2,
				3: 5,
				4: 0,
			},
			want: []int{4, 2, 1, 3},
		},
		{
			input: map[int]float64{},
			want:  []int{},
		},
		{
			input: map[int]float64{
				1: 1,
				2: 2,
				3: 3,
			},
			want: []int{3, 1, 2},
			prioritize: map[int]float64{
				3: 0,
			},
		},
	}

	for _, test := range tests {
		q := NewPriorityQueue[int](0)
		for v, p := range test.input {
			q.Push(v, p)
		}

		if q.Len() != len(test.input) {
			t.Fatalf("%v != %v", q.Len(), len(test.input))
		}

		for v, p := range test.prioritize {
			q.Prioritize(v, p)
		}

		for _, want := range test.want {
			got := q.Pop()
			if !cmp.Equal(*got, want) {
				t.Errorf("%v != %v", *got, want)
			}
		}
	}
}
