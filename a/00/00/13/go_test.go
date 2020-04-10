package a000013

import (
	"testing"

	"github.com/superloach/goeis"
)

func TestA000013(t *testing.T) {
	goeis.TestSeq(
		t, Seq, "0",
		"1", "1", "2", "2", "4",
		"4", "8", "10", "20", "30",
	)
}
