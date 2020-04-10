package a000000

import (
	"testing"

	"github.com/superloach/goeis"
)

func TestA000000(t *testing.T) {
	goeis.TestSequencer(
		t, Seq, 0,
	)
}
