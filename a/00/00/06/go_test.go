package a000006

import (
	"testing"

	"github.com/superloach/goeis"
)

func TestA000006(t *testing.T) {
	goeis.TestSeq(
		t, Seq, "1",
		"1", "1", "2", "2", "3",
		"3", "4", "4", "4", "5",
	)
}
