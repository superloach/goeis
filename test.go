package goeis

import (
	"math/big"
	"testing"
)

func TestSeq(t *testing.T, A Seq, start int, values ...string) {
	a := big.NewInt(0)
	val := big.NewInt(0)

	for i := 0; i < len(values); i++ {
		j := i + start

		a, err := A(j, a)
		if err != nil {
			t.Error(err)
		}

		val, ok := val.SetString(values[i], 10)
		if !ok {
			t.Errorf("val: can't setstring " + values[i])
		}

		if a.Cmp(val) != 0 {
			t.Errorf("A(%d) should be %d, got %d", j, val, a)
		}
	}
}
