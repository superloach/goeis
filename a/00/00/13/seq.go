package a000013

import (
	"github.com/superloach/goeis"
	a000010 "github.com/superloach/goeis/a/00/00/10"
)

var Seq goeis.Seq = func(n int) (int, error) {
	a := 0

	if n < 0 {
		return 0, goeis.ErrOutOfBounds
	}

	if n == 0 {
		return 1, nil
	}

	for _, d := range util.Divides(n) {
		p, err := a000010.Seq(2 * d)
		if err != nil {
			return 0, err
		}

		a += p * (1 << (n / d))
	}

	return a / (2 * n), nil
}
