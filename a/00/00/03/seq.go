package a000003

import (
	"math/big"

	"github.com/superloach/goeis"
)

var Seq goeis.Seq = func(n int, a *big.Int) (*big.Int, error) {
	if n < 1 || n > len(values) {
		return nil, goeis.ErrOutOfBounds
	}

	a, ok := a.SetString(values[n-1], 10)
	if !ok {
		return nil, goeis.ErrSetString
	}

	return a, nil
}
