package a000010

import (
	"testing"

	"github.com/superloach/goeis"
)

func TestA000010(t *testing.T) {
	goeis.TestSeq(
		t, Seq, "1",
		"1", "1", "2", "2", "4",
		"2", "6", "4", "6", "4",
	)
}
