package adt

import (
	"reflect"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	tests := []struct {
		input      map[string]float64
		want       []interface{}
		prioritize map[interface{}]float64
	}{
		{
			input: map[string]float64{
				"a": 3,
				"b": 2,
				"c": 5,
				"d": 0,
			},
			want: []interface{}{"d", "b", "a", "c"},
		},
		{
			input: map[string]float64{},
			want:  []interface{}{},
		},
		{
			input: map[string]float64{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			want: []interface{}{"c", "a", "b"},
			prioritize: map[interface{}]float64{
				"c": 0,
			},
		},
	}

	for _, test := range tests {
		q := NewPriorityQueue(0)
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
			if !reflect.DeepEqual(got, want) {
				t.Errorf("%v != %v", got, want)
			}
		}
	}
}
