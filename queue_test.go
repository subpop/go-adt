package adt

import (
	"testing"
)

func TestQueueString(t *testing.T) {
	tests := []struct {
		input, want []string
	}{
		{
			input: []string{"a", "b", "c"},
			want:  []string{"a", "b", "c"},
		},
	}

	for _, test := range tests {
		var q Queue[string]
		for _, i := range test.input {
			q.Enqueue(i)
		}

		for _, want := range test.want {
			got := q.Dequeue()
			if *got != want {
				t.Errorf("%+v != %+v", got, want)
			}
		}
	}
}

func TestQueueInt(t *testing.T) {
	tests := []struct {
		input, want []int
	}{
		{
			input: []int{1, 2, 3},
			want:  []int{1, 2, 3},
		},
	}

	for _, test := range tests {
		var q Queue[int]
		for _, i := range test.input {
			q.Enqueue(i)
		}

		for _, want := range test.want {
			got := q.Dequeue()
			if *got != want {
				t.Errorf("%+v != %+v", got, want)
			}
		}
	}
}
