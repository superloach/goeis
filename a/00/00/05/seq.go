package a000005

import (
	"math/big"

	"github.com/superloach/goeis"
)

var one = big.NewInt(1)

func Seq(n *big.Int, a *big.Int) (*big.Int, error) {
	a.SetInt64(0)

	s := n.Sign()
	if s == -1 || s == 0 {
		return nil, goeis.ErrOutOfBounds
	}

	t := &big.Int{}

	for x := big.NewInt(1); x.Cmp(n) != 1; x.Add(x, one) {
		for y := big.NewInt(1); y.Cmp(n) != 1; y.Add(y, one) {
			t.Mul(x, y)
			if t.Cmp(n) == 0 {
				a.Add(a, one)
			}
		}
	}

	return a, nil
}
