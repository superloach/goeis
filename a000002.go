package goeis

// Kolakoski sequence: a(n) is length of n-th run; a(1) = 1; sequence consists just of 1's and 2's.
type A000002 struct {
	values []int
	index  int
	iter   int
}

func (a *A000002) Name() string {
	return "A000002"
}

func (a *A000002) gen() {
	var x int
	if a.iter < len(a.values) {
		x = a.values[a.iter]
	} else {
		x = a.iter + 1
	}
	for i := 0; i < x; i++ {
		a.values = append(a.values, 2-((a.iter+1)%2))
	}
	a.iter++
}

func (a *A000002) Next() (n int, ok bool) {
	if a.values == nil {
		a.values = make([]int, 0)
	}

	for len(a.values) <= a.index {
		a.gen()
	}

	n = a.values[a.index]
	ok = true
	a.index++

	return
}

func (a *A000002) Reset() {
	a.index = 0
}
