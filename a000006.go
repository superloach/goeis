package goeis

import "math"

// Integer part of square root of n-th prime.
type A000006 struct {
	primes *A000040
}

func (a *A000006) Name() string {
	return "A000006"
}

func (a *A000006) Next() (n int, ok bool) {
	if a.primes == nil {
		a.primes = &A000040{}
	}

	pn, pok := a.primes.Next()
	if !pok {
		n = 0
		ok = false

		return
	}

	n = int(math.Sqrt(float64(pn)))
	ok = true

	return
}

func (a *A000006) Reset() {
	if a.primes == nil {
		a.primes = &A000040{}
	}

	a.primes.Reset()
}
