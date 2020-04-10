package a000040

import (
	"math/big"

	"github.com/superloach/goeis"
)

var values = []*big.Int{
	big.NewInt(2),
}
var one = big.NewInt(1)

var Seq goeis.Seq = func(n int, a *big.Int) (*big.Int, error) {
	if n < 1 {
		return nil, goeis.ErrOutOfBounds
	}

	i := big.NewInt(0)

	for n > len(values) {
		i.Set(values[len(values)-1])
		i.Add(i, one)

		for {
			if i.ProbablyPrime(10) {
				values = append(values, i)
				break
			}

			i.Add(i, one)
		}
	}

	a.Set(values[n-1])

	return a, nil
}
