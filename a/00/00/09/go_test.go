package a000009

import (
	"testing"

	"github.com/superloach/goeis"
)

func TestA000009(t *testing.T) {
	goeis.TestSequencer(
		t, Seq, 0,
		1, 1, 1, 2, 2,
		3, 4, 5, 6, 8,
	)
}
