package adt

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDisjointSet(t *testing.T) {
	tests := []struct {
		input []*DisjointSet[string]
		want  *DisjointSet[string]
	}{
		{
			input: []*DisjointSet[string]{
				NewDisjointSet("a"),
				NewDisjointSet("b"),
			},
			want: func() *DisjointSet[string] {
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

		if !cmp.Equal(got, test.want, cmp.AllowUnexported(DisjointSet[string]{})) {
			t.Errorf("%+v != %+v", got, test.want)
		}
	}
}
