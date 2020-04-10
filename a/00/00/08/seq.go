package a000008

import (
	"math/big"

	"github.com/superloach/goeis"
)

func Seq(n *big.Int, a *big.Int) (*big.Int, error) {
	if n.Sign() == -1 {
		return nil, goeis.ErrOutOfBounds
	}

	// http://oeis.org/A187243/a187243_1.pdf

	return a, nil
}
