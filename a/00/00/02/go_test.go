package a000002

import (
	"testing"

	"github.com/superloach/goeis"
)

func TestA000002(t *testing.T) {
	goeis.TestSequencer(
		t, Seq, 1,
		1, 2, 2, 1, 1,
		2, 1, 2, 2, 1,
	)
}
