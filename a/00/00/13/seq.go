package a000013

import (
	"math/big"

	"github.com/superloach/goeis"
	a000010 "github.com/superloach/goeis/a/00/00/10"
)

var (
	one = big.NewInt(1)
	two = big.NewInt(2)
)

func Seq(n *big.Int, a *big.Int) (*big.Int, error) {
	s := n.Sign()
	if s == -1 {
		return nil, goeis.ErrOutOfBounds
	}
	if s == 0 {
		return one, nil
	}

	a.SetInt64(0)

	t := &big.Int{}
	d := (&big.Int{}).Set(one)

	// d in 1..n where gcd(d, n) == d
	for {
		if d.Cmp(n) == 1 {
			break
		}

		if t.GCD(nil, nil, d, n).Cmp(d) != 0 {
			d.Add(d, one)
			continue
		}

		e := (&big.Int{}).Mul(d, two)

		p := &big.Int{}
		p, err := a000010.Seq(e, p)
		if err != nil {
			return nil, err
		}

		f := (&big.Int{})
		f.Quo(n, d)
		f.Exp(two, f, nil)
		f.Mul(p, f)

		a.Add(a, f)

		d.Add(d, one)
	}

	t.SetInt64(2)
	t.Mul(t, n)
	a.Quo(a, t)

	return a, nil
}
