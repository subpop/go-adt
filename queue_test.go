package ds

import (
	"testing"
)

func TestQueue(t *testing.T) {
	tests := []struct {
		input, want []interface{}
	}{
		{
			input: []interface{}{"a", "b", "c"},
			want:  []interface{}{"a", "b", "c"},
		},
	}

	for _, test := range tests {
		var q Queue
		for _, i := range test.input {
			q.Enqueue(i)
		}

		for _, want := range test.want {
			got := q.Dequeue()
			if got != want {
				t.Errorf("%+v != %+v", got, want)
			}
		}
	}
}
