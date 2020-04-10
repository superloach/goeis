package a000003

import "github.com/superloach/goeis"

var values = []int{
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

var Seq goeis.Seq = func(n int) (int, error) {
	// 1-indexed
	n--

	if n < 0 || n >= len(values) {
		return 0, goeis.ErrOutOfBounds
	}

	return values[n], nil
}