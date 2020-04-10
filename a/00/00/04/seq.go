package a000004

import (
	"math/big"
)

func Seq(_ *big.Int, a *big.Int) (*big.Int, error) {
	a.SetInt64(0)

	return a, nil
}
