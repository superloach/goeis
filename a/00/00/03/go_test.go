package a000003

import (
	"testing"

	"github.com/superloach/goeis"
)

func TestA000003(t *testing.T) {
	goeis.TestSequencer(
		t, Seq, 1,
		1, 1, 1, 1, 2,
		2, 1, 2, 2, 2,
	)
}
