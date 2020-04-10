package a000008

import "math"

var Seq = func(n int) (int, error) {
	// http://oeis.org/A187243/a187243_1.pdf
	q := n / 10
	a := float64(q+1) * (math.Round(float64((n+4)*(n+4))/20) - ((1 / float64(6)) * float64(q) * float64(3*n+10*q+7)))

	return int(a), nil
}
