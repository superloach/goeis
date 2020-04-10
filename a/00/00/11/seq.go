package a000011

import (
	"math/big"

	"github.com/superloach/goeis"
	a000013 "github.com/superloach/goeis/a/00/00/13"
)

var two = big.NewInt(2)

func Seq(n *big.Int, a *big.Int) (*big.Int, error) {
	if n.Sign() == -1 {
		return nil, goeis.ErrOutOfBounds
	}

	a, err := a000013.Seq(n, a)
	if err != nil {
		return nil, err
	}

	t := (&big.Int{}).Set(n)
	t.Quo(t, two)
	t.Exp(two, t, nil)

	a.Add(a, t)
	a.Quo(a, two)

	return a, nil
}
