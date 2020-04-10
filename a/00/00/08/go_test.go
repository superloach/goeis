package a000008

import (
	"testing"

	"github.com/superloach/goeis"
)

func TestA000008(t *testing.T) {
	goeis.TestSeq(
		t, Seq, "0",
		"1", "1", "2", "2", "3",
		"4", "5", "6", "7", "8",
	)
}
