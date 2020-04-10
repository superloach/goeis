package a000001

import (
	"testing"

	"github.com/superloach/goeis"
)

func TestA000001(t *testing.T) {
	goeis.TestSeq(
		t, Seq, "0",
		"0", "1", "1", "1", "2",
		"1", "2", "1", "5", "2",
	)
}
