package adt

import (
	"reflect"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	tests := []struct {
		input map[string]int
		want  []interface{}
	}{
		{
			input: map[string]int{
				"a": 3,
				"b": 2,
				"c": 5,
				"d": 0,
			},
			want: []interface{}{"d", "b", "a", "c"},
		},
		{
			input: map[string]int{},
			want:  []interface{}{},
		},
	}

	for _, test := range tests {
		q := NewPriorityQueue(0)
		for v, p := range test.input {
			q.Push(v, p)
		}

		for _, want := range test.want {
			got := q.Pop()
			if !reflect.DeepEqual(got, want) {
				t.Errorf("%v != %v", got, want)
			}
		}
	}
}
