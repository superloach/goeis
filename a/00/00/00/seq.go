package a000000

import (
	"math/big"

	"github.com/superloach/goeis"
)

var Seq goeis.Seq = func(_ int, a *big.Int) (*big.Int, error) {
	return nil, goeis.ErrOutOfBounds
}
