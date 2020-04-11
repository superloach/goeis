package a000008

import (
	"math/big"

	"github.com/superloach/goeis"
)

var half = big.NewFloat(0.5)
var one = big.NewInt(1)
var onef = big.NewFloat(1)
var three = big.NewInt(3)
var four = big.NewInt(4)
var six = big.NewFloat(6)
var seven = big.NewInt(7)
var ten = big.NewInt(10)
var twenty = big.NewFloat(20)

func Seq(n *big.Int, a *big.Int) (*big.Int, error) {
	if n.Sign() == -1 {
		return nil, goeis.ErrOutOfBounds
	}

	// http://oeis.org/A187243/a187243_1.pdf
	q := (&big.Int{}).Quo(n, ten)

	b := (&big.Int{}).Add(n, four)
	c := (&big.Int{}).Mul(b, b)
	d := (&big.Float{}).SetInt(c)
	e := (&big.Float{}).Quo(d, twenty)
	e.Add(e, half)
	f := &big.Int{}
	f, _ = e.Int(f)
	g := (&big.Int{}).Mul(n, three)
	h := (&big.Int{}).Mul(q, ten)
	i := (&big.Int{}).Add(g, h)
	j := (&big.Int{}).Add(i, seven)
	k := (&big.Float{}).SetInt(j)
	l := (&big.Float{}).SetInt(q)
	l.Quo(l, six)
	m := (&big.Float{}).Mul(k, l)
	o := (&big.Float{}).SetInt(f)
	p := (&big.Float{}).Add(m, o)
	r := (&big.Float{}).SetInt(q)
	s := (&big.Float{}).Add(r, onef)
	t := (&big.Float{}).Mul(s, p)
	a, _ = t.Int(a)

	return a, nil
}
