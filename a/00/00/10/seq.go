package a000010

import (
	"github.com/superloach/goeis"
	"github.com/superloach/goeis/util"
)

var Seq goeis.Seq = func(n int) (int, error) {
	if n < 1 {
		return 0, goeis.ErrOutOfBounds
	}

	a := 0

	for i := 1; i <= n; i++ {
		if util.GCD(i, n) == 1 {
			a++
		}
	}

	return a, nil
}
