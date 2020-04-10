package a000005

import (
	"testing"

	"github.com/superloach/goeis"
)

func TestA000005(t *testing.T) {
	goeis.TestSeq(
		t, Seq, 1,
		"1", "2", "2", "3", "2",
		"4", "2", "4", "3", "4",
	)
}
