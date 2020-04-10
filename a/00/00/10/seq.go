package a000010

import (
	"math/big"

	"github.com/superloach/goeis"
)

var one = big.NewInt(1)

var Seq goeis.Seq = func(n int, a *big.Int) (*big.Int, error) {
	if n < 1 {
		return nil, goeis.ErrOutOfBounds
	}

	a.SetInt64(0)

	m := big.NewInt(int64(n))
	i := &big.Int{}

	for j := int64(1); j <= int64(n); j++ {
		i.SetInt64(j)

		i.GCD(nil, nil, i, m)

		if i.Cmp(one) == 0 {
			a.Add(a, one)
		}
	}

	return a, nil
}
