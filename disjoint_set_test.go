package adt

import (
	"reflect"
	"testing"
)

func TestDisjointSet(t *testing.T) {
	tests := []struct {
		input []*DisjointSet
		want  *DisjointSet
	}{
		{
			input: []*DisjointSet{
				NewDisjointSet("a"),
				NewDisjointSet("b"),
			},
			want: func() *DisjointSet {
				c := NewDisjointSet("c")
				c.size = 3
				return c
			}(),
		},
	}

	for _, test := range tests {
		got := NewDisjointSet("c")
		for _, s := range test.input {
			got = Union(got, s)
		}

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%+v != %+v", got, test.want)
		}
	}
}
