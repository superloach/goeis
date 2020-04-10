package a000004

import (
	"math/big"

	"github.com/superloach/goeis"
)

var Seq goeis.Seq = func(_ int, a *big.Int) (*big.Int, error) {
	return a.SetInt64(0), nil
}
