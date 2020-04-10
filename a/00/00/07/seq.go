package a000007

import "github.com/superloach/goeis"

var Seq goeis.Seq = func(n int) (int, error) {
	if n == 0 {
		return 1, nil
	}

	return 0, nil
}
