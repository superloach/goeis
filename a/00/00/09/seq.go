package a000009

import "github.com/superloach/goeis"

var Seq goeis.Seq = func(n int) (int, error) {
	if n < 0 || n >= len(values) {
		return 0, goeis.ErrOutOfBounds
	}

	return values[n], nil
}
