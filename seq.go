package goeis

import "math/big"

type Seq func(int, *big.Int) (*big.Int, error)
