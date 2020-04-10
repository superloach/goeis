package goeis

import (
	"math/big"
	"testing"
)

var one = big.NewInt(1)

func TestSeq(t *testing.T, A Seq, start string, vals ...string) {
	n, ok := (&big.Int{}).SetString(start, 10)
	if !ok {
		t.Fatalf("n: can't setstring " + start)
	}

	a := &big.Int{}
	v := &big.Int{}

	for i := 0; i < len(vals); i++ {
		a, err := A(n, a)
		if err != nil {
			t.Fatal(err)
		}

		v, ok := v.SetString(vals[i], 10)
		if !ok {
			t.Fatalf("v: can't setstring " + vals[i])
		}

		if a.Cmp(v) != 0 {
			t.Fatalf("A(%d) should be %d, got %d", n, v, a)
		}

		n.Add(n, one)
	}
}
