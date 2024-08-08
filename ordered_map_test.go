package adt

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestOrderedMapStringAppend(t *testing.T) {
	tests := []struct {
		description string
		have        *OrderedMap[string, string]
		input       struct{ k, v string }
		want        *OrderedMap[string, string]
		wantError   error
	}{
		{
			have: &OrderedMap[string, string]{
				order:  []string{"a"},
				values: map[string]string{"a": "0"},
				len:    1,
			},
			input: struct {
				k string
				v string
			}{
				k: "b",
				v: "1",
			},
			want: &OrderedMap[string, string]{
				order: []string{"a", "b"},
				values: map[string]string{
					"a": "0",
					"b": "1",
				},
				len: 2,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			err := test.have.Append(test.input.k, test.input.v)

			if test.wantError != nil {
				if !cmp.Equal(test.wantError, err, cmpopts.EquateErrors()) {
					t.Errorf("wantError: %v - got: %v", test.wantError, err)
				}
			} else {
				if !cmp.Equal(test.have, test.want, cmp.AllowUnexported(OrderedMap[string, string]{}), cmpopts.IgnoreFields(OrderedMap[string, string]{}, "mu")) {
					t.Errorf("%v != %v", test.have, test.want)
				}
			}
		})
	}
}

func TestOrderedMapStringDelete(t *testing.T) {
	tests := []struct {
		description string
		have        *OrderedMap[string, string]
		input       string
		want        *OrderedMap[string, string]
	}{
		{
			have: &OrderedMap[string, string]{
				order: []string{"a", "b"},
				values: map[string]string{
					"a": "0",
					"b": "1",
				},
				len: 2,
			},
			input: "a",
			want: &OrderedMap[string, string]{
				order:  []string{"b"},
				values: map[string]string{"b": "1"},
				len:    1,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			test.have.Delete(test.input)

			if !cmp.Equal(test.have, test.want, cmp.AllowUnexported(OrderedMap[string, string]{}), cmpopts.IgnoreFields(OrderedMap[string, string]{}, "mu")) {
				t.Errorf("%v != %v", test.have, test.want)
			}
		})
	}
}

func TestOrderedMapStringVisit(t *testing.T) {
	type element struct{ k, v string }
	tests := []struct {
		description string
		have        *OrderedMap[string, string]
		want        []element
	}{
		{
			have: &OrderedMap[string, string]{
				order: []string{"a", "b", "c"},
				values: map[string]string{
					"a": "0",
					"b": "1",
					"c": "2",
				},
				len: 3,
			},
			want: []element{
				{"a", "0"},
				{"b", "1"},
				{"c", "2"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var got []element
			test.have.Visit(func(key, value string) {
				got = append(got, element{key, value})
			})

			if !cmp.Equal(got, test.want, cmp.AllowUnexported(OrderedMap[string, string]{}), cmpopts.IgnoreFields(OrderedMap[string, string]{}, "mu"), cmpopts.EquateComparable(element{})) {
				t.Errorf("%v", cmp.Diff(got, test.want))
			}
		})
	}
}
