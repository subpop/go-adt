package adt

import (
	"testing"

	"github.com/google/go-cmp/cmp"
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

		if !cmp.Equal(got, test.want, cmp.AllowUnexported(DisjointSet{})) {
			t.Errorf("%+v != %+v", got, test.want)
		}
	}
}
