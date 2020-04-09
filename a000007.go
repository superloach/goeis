package goeis

// The characteristic function of {0}: a(n) = 0^n.
type A000007 struct {
	index  int
}

func (a *A000007) Name() string {
	return "A000007"
}

func (a *A000007) Next() (n int, ok bool) {
	if a.index == 0 {
		n = 1
	}
	ok = true
	a.index++

	return
}

func (a *A000007) Reset() {
	a.index = 0
}
