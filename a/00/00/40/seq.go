package a000040

import (
	"math"

	"github.com/superloach/goeis"
)

var values = []int{2}

var Seq goeis.Seq = func(n int) (int, error) {
	n--

	if n < 0 {
		return 0, goeis.ErrOutOfBounds
	}

	for len(values) <= n {
		for i := values[len(values)-1] + 1; ; i++ {
			prime := true

			for x := 2; x <= int(math.Sqrt(float64(i))); x++ {
				if i%x == 0 {
					prime = false
				}
			}

			if prime {
				values = append(values, i)
				break
			}
		}
	}

	return values[n], nil
}
