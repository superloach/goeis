package a000040

import (
	"math/big"

	"github.com/superloach/goeis"
)

var values = map[string]*big.Int{
	"0": big.NewInt(2),
}
var count = big.NewInt(1)

var one = big.NewInt(1)

func Seq(n *big.Int, a *big.Int) (*big.Int, error) {
	s := n.Sign()
	if s == -1 || s == 0 {
		return nil, goeis.ErrOutOfBounds
	}

	i := big.NewInt(0)

	for n.Cmp(count) == 1 {
		x := (&big.Int{}).Sub(count, one)
		i.Set(values[x.String()])
		i.Add(i, one)

		for {
			if i.ProbablyPrime(10) {
				values[count.String()] = i
				count.Add(count, one)

				break
			}

			i.Add(i, one)
		}
	}

	m := (&big.Int{}).Sub(n, one)

	a.Set(values[m.String()])

	return a, nil
}
