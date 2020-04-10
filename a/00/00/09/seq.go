package a000009

import (
	"math/big"

	"github.com/superloach/goeis"
)

func Seq(n *big.Int, a *big.Int) (*big.Int, error) {
	if v, ok := values[n.String()]; ok {
		a.Set(v)
		return a, nil
	} else {
		return nil, goeis.ErrOutOfBounds
	}
}
