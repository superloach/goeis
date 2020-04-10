package a000006

import (
	"math/big"

	"github.com/superloach/goeis"
	a000040 "github.com/superloach/goeis/a/00/00/40"
)

var Seq goeis.Seq = func(n int, a *big.Int) (*big.Int, error) {
	if n < 1 {
		return nil, goeis.ErrOutOfBounds
	}

	a, err := a000040.Seq(n, a)
	if err != nil {
		return nil, err
	}

	a.Sqrt(a)

	return a, nil
}
