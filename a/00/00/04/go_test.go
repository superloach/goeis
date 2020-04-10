package a000004

import (
	"testing"

	"github.com/superloach/goeis"
)

func TestA000004(t *testing.T) {
	goeis.TestSeq(
		t, Seq, 0,
		"0", "0", "0", "0", "0",
		"0", "0", "0", "0", "0",
	)
}
