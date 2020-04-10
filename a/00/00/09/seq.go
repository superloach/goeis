package a000009

import "github.com/superloach/goeis"

var values = []int{
	1, 1, 1, 2, 2, 3, 4, 5, 6, 8,
	10, 12, 15, 18, 22, 27, 32, 38, 46, 54,
	64, 76, 89, 104, 122, 142, 165, 192, 222, 256,
	296, 340, 390, 448, 512, 585, 668, 760, 864, 982,
	1113, 1260, 1426, 1610, 1816, 2048, 2304, 2590, 2910, 3264,
	3658, 4097, 4582, 5120, 5718, 6378,
}

var Seq goeis.Seq = func(n int) (int, error) {
	if n < 0 || n >= len(values) {
		return 0, goeis.ErrOutOfBounds
	}

	return values[n], nil
}
