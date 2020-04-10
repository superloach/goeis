package a000012

import "math/big"

func Seq(_ *big.Int, a *big.Int) (*big.Int, error) {
	a.SetInt64(1)

	return a, nil
}
