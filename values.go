package goeis

import "math/big"

func Values(orig map[string]string) map[string]*big.Int {
	conv := make(map[string]*big.Int)

	for k, v := range orig {
		vi, ok := (&big.Int{}).SetString(v, 10)
		if !ok {
			panic(ErrSetString)
		}

		conv[k] = vi
	}

	return conv
}
