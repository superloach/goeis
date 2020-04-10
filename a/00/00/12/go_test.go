package a000012

import (
	"testing"

	"github.com/superloach/goeis"
)

func TestA000012(t *testing.T) {
	goeis.TestSeq(
		t, Seq, "0",
		"1", "1", "1", "1", "1",
		"1", "1", "1", "1", "1",
	)
}
