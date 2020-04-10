package goeis

import "math/big"

type Seq func(*big.Int, *big.Int) (*big.Int, error)
