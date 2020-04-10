package a000011

import (
	"testing"

	"github.com/superloach/goeis"
)

func TestA000011(t *testing.T) {
	goeis.TestSequencer(
		t, Seq, 0,
		1, 1, 2, 2, 4,
		4, 8, 9, 18, 23,
	)
}
