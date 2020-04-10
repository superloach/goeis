package a000007

import "math/big"

func Seq(n *big.Int, a *big.Int) (*big.Int, error) {
	if n.Sign() == 0 {
		a.SetInt64(1)
	} else {
		a.SetInt64(0)
	}

	return a, nil
}
