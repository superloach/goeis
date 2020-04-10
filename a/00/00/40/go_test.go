package a000040

import (
	"testing"

	"github.com/superloach/goeis"
)

func TestA000040(t *testing.T) {
	goeis.TestSeq(
		t, Seq, "1",
		"2", "3", "5", "7", "11",
		"13", "17", "19", "23", "29",
	)
}
