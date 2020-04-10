package a000005

import (
	"math/big"

	"github.com/superloach/goeis"
)

var one = big.NewInt(1)

var Seq goeis.Seq = func(n int, a *big.Int) (*big.Int, error) {
	a.Sub(a, a)

	if n < 1 {
		return nil, goeis.ErrOutOfBounds
	}

	for x := 1; x <= n; x++ {
		for y := 1; y <= n; y++ {
			if x*y == n {
				a.Add(a, one)
			}
		}
	}

	return a, nil
}
