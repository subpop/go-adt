package adt

import (
	"testing"
)

func TestStack(t *testing.T) {
	tests := []struct {
		input, want []interface{}
	}{
		{
			input: []interface{}{"a", "b", "c"},
			want:  []interface{}{"c", "b", "a"},
		},
	}

	for _, test := range tests {
		var s Stack
		for _, i := range test.input {
			s.Push(i)
		}

		for _, want := range test.want {
			got := s.Pop()
			if got != want {
				t.Errorf("%+v != %+v", got, want)
			}
		}
	}
}
