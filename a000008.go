package goeis

// Number of ways of making change for n cents using coins of 1, 2, 5, 10 cents.
type A000008 struct {
	cache map[int]int
	index int
}

func (a *A000008) Name() string {
	return "A000008"
}

func (a *A000008) Next() (n int, ok bool) {
	if a.cache == nil {
		a.cache = make(map[int]int)
		a.cache[0] = 1
	}

	for _, coin := range []int{1, 2, 5, 10} {
		for i := coin; i <= a.index + 1; i++ {
			b, _ := a.cache[i]
			c, _ := a.cache[i - coin]
			a.cache[i] = b + c
		}
	}

	n = a.cache[a.index + 1]
	ok = true
	a.index++

	return
}

func (a *A000008) Reset() {
	a.index = 0
}
