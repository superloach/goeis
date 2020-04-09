package goeis

// d(n) (also called tau(n) or sigma_0(n)), the number of divisors of n.
type A000005 struct {
	index  int
}

func (a *A000005) Name() string {
	return "A000005"
}

func (a *A000005) Next() (n int, ok bool) {
	for x := 1; x <= a.index + 1; x++ {
		for y := 1; y <= a.index + 1; y++ {
			if x*y == a.index + 1 {
				n++
			}
		}
	}

	ok = true
	a.index++

	return
}

func (a *A000005) Reset() {
	a.index = 0
}
