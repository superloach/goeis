package a000011

import (
	"github.com/superloach/goeis"
	a000013 "github.com/superloach/goeis/a/00/00/13"
)

var Seq goeis.Seq = func(n int) (int, error) {
	if n < 0 {
		return 0, goeis.ErrOutOfBounds
	}

	a, err := a000013.Seq(n)
	if err != nil {
		return 0, err
	}

	return (a + (1 << (n / 2))) / 2, nil
}
