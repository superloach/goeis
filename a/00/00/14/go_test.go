package a000014

import (
	"testing"

	"github.com/superloach/goeis"
)

func TestA000014(t *testing.T) {
	goeis.TestSeq(
		t, Seq, "0",
		"0", "1", "1", "0", "1",
		"1", "2", "2", "4", "5",
	)
}
