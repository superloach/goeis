package a000002

import "github.com/superloach/goeis"

var values = make([]int, 0)
var iter = 0

var Seq goeis.Seq = func(n int) (int, error) {
	if n < 1 {
		return 0, goeis.ErrOutOfBounds
	}

	for n > len(values) {
		x := 0
		if iter < len(values) {
			x = values[iter]
		} else {
			x = iter + 1
		}

		for i := 0; i < x; i++ {
			values = append(values, 2-((iter+1)%2))
		}
		iter++
	}

	return values[n-1], nil
}
