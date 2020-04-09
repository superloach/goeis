package goeis

var a000001 = []int{
	0, 1, 1, 1, 2, 1, 2, 1, 5, 2, 2, 1, 5, 1, 2, 1, 14, 1, 5,
	1, 5, 2, 2, 1, 15, 2, 2, 5, 4, 1, 4, 1, 51, 1, 2, 1, 14, 1,
	2, 2, 14, 1, 6, 1, 4, 2, 2, 1, 52, 2, 5, 1, 5, 1, 15, 2, 13,
	2, 2, 1, 13, 1, 2, 4, 267, 1, 4, 1, 5, 1, 4, 1, 50, 1, 2, 3,
	4, 1, 6, 1, 52, 15, 2, 1, 15, 1, 2, 1, 12, 1, 10, 1, 4, 2,
}

// Number of groups of order n.
// TODO: proper generation
type A000001 struct {
	index int
}

func (a *A000001) Name() string {
	return "A000001"
}

func (a *A000001) Next() (n int, ok bool) {
	if a.index >= len(a000001) {
		return
	}

	n = a000001[a.index]
	ok = true
	a.index++

	return
}

func (a *A000001) Reset() {
	a.index = 0
}
