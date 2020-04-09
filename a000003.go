package goeis

var a000003 = []int{
	1, 1, 1, 1, 2, 2, 1, 2, 2, 2,
	3, 2, 2, 4, 2, 2, 4, 2, 3, 4,
	4, 2, 3, 4, 2, 6, 3, 2, 6, 4,
	3, 4, 4, 4, 6, 4, 2, 6, 4, 4,
	8, 4, 3, 6, 4, 4, 5, 4, 4, 6,
	6, 4, 6, 6, 4, 8, 4, 2, 9, 4,
	6, 8, 4, 4, 8, 8, 3, 8, 8, 4,
	7, 4, 4, 10, 6, 6, 8, 4, 5, 8,
	6, 4, 9, 8, 4, 10, 6, 4, 12, 8,
	6, 6, 4, 8, 8, 8, 4, 8, 6, 4,
}

// Number of classes of primitive positive definite binary quadratic forms of discriminant D = -4n; or equivalently the class number of the quadratic order of discriminant D = -4n.
// TODO: proper generation
type A000003 struct {
	index int
}

func (a *A000003) Name() string {
	return "A000003"
}

func (a *A000003) Next() (n int, ok bool) {
	if a.index >= len(a000003) {
		return
	}

	n = a000003[a.index]
	ok = true
	a.index++

	return
}

func (a *A000003) Reset() {
	a.index = 0
}
