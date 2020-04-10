package a000005

import "github.com/superloach/goeis"

var Seq goeis.Seq = func(n int) (int, error) {
	a := 0

	if n < 1 {
		return 0, goeis.ErrOutOfBounds
	}

	for x := 1; x <= n; x++ {
		for y := 1; y <= n; y++ {
			if x*y == n {
				a++
			}
		}
	}

	return a, nil
}
