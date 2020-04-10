package a000040

import (
	"github.com/superloach/goeis"
	"github.com/superloach/goeis/util"
)

var values = []int{2}

var Seq goeis.Seq = func(n int) (int, error) {
	n--

	if n < 0 {
		return 0, goeis.ErrOutOfBounds
	}

	for len(values) <= n {
		for i := values[len(values)-1] + 1; ; i++ {
			if util.IsPrime(i) {
				values = append(values, i)
				break
			}
		}
	}

	return values[n], nil
}
