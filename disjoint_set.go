package adt

// DisjointSet is a data structure that tracks a set of elements partitioned
// into a number of non-overlapping (disjoint) subsets.
type DisjointSet struct {
	Value  interface{}
	parent *DisjointSet
	size   int
}

// NewDisjointSet returns a disjoint-set initialized to contain only value v.
func NewDisjointSet(v interface{}) *DisjointSet {
	d := &DisjointSet{
		Value: v,
		size:  1,
	}
	d.parent = d
	return d
}

// Find finds the root (or representative element) of disjoint-set d.
func (d *DisjointSet) Find() *DisjointSet {
	if d.parent != d {
		d.parent = d.parent.Find()
	}
	return d.parent
}

// Union finds thes the representative element of x and y and merges the
// smaller of the two sets into the other.
func Union(x, y *DisjointSet) *DisjointSet {
	xRoot := x.Find()
	yRoot := y.Find()

	// x and y are already in the same set
	if xRoot == yRoot {
		return xRoot
	}

	// x and y are not in same set, so we merge them
	if xRoot.size < yRoot.size {
		xRoot, yRoot = yRoot, xRoot // swap xRoot and yRoot
	}

	// merge yRoot into xRoot
	yRoot.parent = xRoot
	xRoot.size += yRoot.size

	return xRoot
}
