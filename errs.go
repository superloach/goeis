package goeis

import "errors"

var (
	ErrMissingGet  = errors.New("missing get")
	ErrOutOfBounds = errors.New("out of bounds")
)
