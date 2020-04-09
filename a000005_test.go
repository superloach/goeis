package goeis

import (
	"strconv"
	"testing"
)

var testA000005 = []int{
	1, 2, 2, 3, 2,
	4, 2, 4, 3, 4,
}

func TestA000005Name(t *testing.T) {
	a := &A000005{}

	name := a.Name()
	if name != "A000005" {
		t.Error("name should be A000005, got " + name)
	}
}

func TestA000005Next(t *testing.T) {
	a := &A000005{}

	for i := 0; i < 10; i++ {
		x, ok := a.Next()
		if !ok {
			t.Error("not ok at " + strconv.Itoa(i))
		}

		if x != testA000005[i] {
			t.Error("x should be " + strconv.Itoa(testA000005[i]) + ", got " + strconv.Itoa(x))
		}
	}
}

func TestA000005Reset(t *testing.T) {
	a := &A000005{}

	for i := 0; i < 10; i++ {
		_, ok := a.Next()
		if !ok {
			t.Error("not ok at " + strconv.Itoa(i))
		}
	}

	a.Reset()

	x, ok := a.Next()
	if !ok {
		t.Error("not ok at 0")
	}

	if x != testA000005[0] {
		t.Error("x should be " + strconv.Itoa(testA000005[0]) + ", got " + strconv.Itoa(x))
	}
}
