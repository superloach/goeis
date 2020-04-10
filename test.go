package goeis

import "testing"

func TestSequencer(t *testing.T, A Seq, start int, values ...int) {
	for i := 0; i < len(values); i++ {
		j := start + i

		a, err := A(j)
		if err != nil {
			t.Error(err)
		}

		if a != values[i] {
			t.Errorf("A(%d) should be %d, got %d", j, values[i], a)
		}
	}
}
