package adt

import (
	"testing"
)

func TestStackString(t *testing.T) {
	tests := []struct {
		input, want []string
	}{
		{
			input: []string{"a", "b", "c"},
			want:  []string{"c", "b", "a"},
		},
	}

	for _, test := range tests {
		var s Stack[string]
		for _, i := range test.input {
			s.Push(i)
		}

		for _, want := range test.want {
			got := s.Pop()
			if *got != want {
				t.Errorf("%+v != %+v", got, want)
			}
		}
	}
}

func TestStackInt(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{
			input: []int{1, 2, 3},
			want:  []int{3, 2, 1},
		},
	}

	for _, test := range tests {
		var s Stack[int]
		for _, i := range test.input {
			s.Push(i)
		}

		for _, want := range test.want {
			got := s.Pop()
			if *got != want {
				t.Errorf("%+v != %+v", got, want)
			}
		}
	}
}
