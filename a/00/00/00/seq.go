package a000000

import "github.com/superloach/goeis"

var Seq goeis.Seq = func(n int) (int, error) {
	return 0, goeis.ErrOutOfBounds
}
