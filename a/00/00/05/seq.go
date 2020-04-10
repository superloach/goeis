package a000005

import "github.com/superloach/goeis"

var Seq goeis.Seq = func(n int) (int, error) {
	n--

	a := 0
	for x := 1; x <= n+1; x++ {
		for y := 1; y <= n+1; y++ {
			if x*y == n+1 {
				a++
			}
		}
	}

	return a, nil
}
