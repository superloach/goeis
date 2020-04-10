package a000007

import (
	"testing"

	"github.com/superloach/goeis"
)

func TestA000007(t *testing.T) {
	goeis.TestSequencer(
		t, Seq, 0,
		1, 0, 0, 0, 0,
		0, 0, 0, 0, 0,
	)
}
