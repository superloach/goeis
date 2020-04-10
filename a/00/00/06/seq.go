package a000006

import (
	"math/big"

	"github.com/superloach/goeis"
	a000040 "github.com/superloach/goeis/a/00/00/40"
)

func Seq(n *big.Int, a *big.Int) (*big.Int, error) {
	s := n.Sign()
	if s == -1 || s == 0 {
		return nil, goeis.ErrOutOfBounds
	}

	a, err := a000040.Seq(n, a)
	if err != nil {
		return nil, err
	}

	a.Sqrt(a)

	return a, nil
}
