package goeis

import "math"

// The prime numbers.
type A000040 struct {
	values []int
	index  int
}

func (a *A000040) Name() string {
	return "A000040"
}

func (a *A000040) gen() {
	if len(a.values) == 0 {
		a.values = append(a.values, 2)
	}

	for i := a.values[len(a.values) - 1] + 1; ; i++ {
		prime := true
		for x := 2; x <= int(math.Sqrt(float64(i))); x++ {
			if i%x == 0 {
				prime = false
			}
		}
		if prime {
			a.values = append(a.values, i)
			return
		}
	}
}

func (a *A000040) Next() (n int, ok bool) {
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

func (a *A000040) Reset() {
	a.index = 0
}
