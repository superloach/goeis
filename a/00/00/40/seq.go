package a000040

import (
	"github.com/superloach/goeis"
	"github.com/superloach/goeis/util"
)

var values = []int{2}

var Seq goeis.Seq = func(n int) (int, error) {
	if n < 1 {
		return 0, goeis.ErrOutOfBounds
	}

	for n > len(values) {
		for i := values[len(values)-1] + 1; ; i++ {
			if util.IsPrime(i) {
				values = append(values, i)
				break
			}
		}
	}

	return values[n-1], nil
}
