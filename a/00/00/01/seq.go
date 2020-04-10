package a000001

import "github.com/superloach/goeis"

var values = []int{
	0, 1, 1, 1, 2, 1, 2, 1, 5, 2,
	2, 1, 5, 1, 2, 1, 14, 1, 5, 1,
	5, 2, 2, 1, 15, 2, 2, 5, 4, 1,
	4, 1, 51, 1, 2, 1, 14, 1, 2, 2,
	14, 1, 6, 1, 4, 2, 2, 1, 52, 2,
	5, 1, 5, 1, 15, 2, 13, 2, 2, 1,
	13, 1, 2, 4, 267, 1, 4, 1, 5, 1,
	4, 1, 50, 1, 2, 3, 4, 1, 6, 1,
	52, 15, 2, 1, 15, 1, 2, 1, 12, 1,
	10, 1, 4, 2,
}

var Seq goeis.Seq = func(n int) (int, error) {
	if n < 0 || n >= len(values) {
		return 0, goeis.ErrOutOfBounds
	}

	return values[n], nil
}
