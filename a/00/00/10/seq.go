package a000010

import (
	"math/big"

	"github.com/superloach/goeis"
)

var one = big.NewInt(1)

func Seq(n *big.Int, a *big.Int) (*big.Int, error) {
	s := n.Sign()
	if s == -1 || s == 0 {
		return nil, goeis.ErrOutOfBounds
	}

	a.SetInt64(0)

	t := &big.Int{}

	for j := big.NewInt(1); j.Cmp(n) != 1; j.Add(j, one) {
		t.GCD(nil, nil, j, n)

		if t.Cmp(one) == 0 {
			a.Add(a, one)
		}
	}

	return a, nil
}
