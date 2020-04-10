package a000002

import (
	"math/big"

	"github.com/superloach/goeis"
)

var values = make(map[string]*big.Int)
var count = &big.Int{}
var iter = &big.Int{}

var one = big.NewInt(1)
var two = big.NewInt(2)

func Seq(n *big.Int, a *big.Int) (*big.Int, error) {
	s := n.Sign()
	if s == -1 || s == 0 {
		return nil, goeis.ErrOutOfBounds
	}

	for n.Cmp(big.NewInt(int64(len(values)))) == 1 {
		x, ok := values[iter.String()]
		if !ok {
			x = (&big.Int{}).Add(iter, one)
		}

		for i := int64(0); i < x.Int64(); i++ {
			// 2 - ((iter + 1) % 2)
			val := big.NewInt(2)
			val.Sub(
				val,
				(&big.Int{}).Mod(
					(&big.Int{}).Add(iter, one),
					two,
				),
			)

			values[count.String()] = val
			count.Add(count, one)
		}

		iter.Add(iter, one)
	}

	m := (&big.Int{}).Sub(n, one)

	a.Set(values[m.String()])

	print(n.String())
	print(": ")
	println(a.String())

	return a, nil
}
