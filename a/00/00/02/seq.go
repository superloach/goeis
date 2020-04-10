package a000002

import (
	"math/big"

	"github.com/superloach/goeis"
)

var values = make([]*big.Int, 0)
var iter = 0

var one = big.NewInt(1)
var two = big.NewInt(2)

var Seq goeis.Seq = func(n int, a *big.Int) (*big.Int, error) {
	if n < 1 {
		return nil, goeis.ErrOutOfBounds
	}

	x := (&big.Int{})

	for n > len(values) {
		x.SetInt64(0)

		if iter < len(values) {
			x.Set(values[iter])
		} else {
			x.Set(big.NewInt(int64(iter) + 1))
		}

		for i := int64(0); i < x.Int64(); i++ {
			// 2 - ((iter + 1) % 2)
			val := (&big.Int{}).Sub(
				big.NewInt(2),
				(&big.Int{}).Mod(
					(&big.Int{}).Add(
						big.NewInt(int64(iter)),
						one,
					),
					two,
				),
			)

			values = append(values, val)
		}

		iter++
	}

	return values[n-1], nil
}
