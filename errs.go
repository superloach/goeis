package goeis

import "errors"

var (
	ErrOutOfBounds = errors.New("out of bounds")
	ErrSetString   = errors.New("set string failed")
)
