package a000006

import (
	"math"

	"github.com/superloach/goeis"
	a000040 "github.com/superloach/goeis/a/00/00/40"
)

var Seq goeis.Seq = func(n int) (int, error) {
	p, err := a000040.Seq(n)
	if err != nil {
		return 0, err
	}

	return int(math.Sqrt(float64(p))), nil
}
