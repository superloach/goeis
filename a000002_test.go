package goeis

import (
	"strconv"
	"testing"
)

var testA000002 = []int{
	1, 2, 2, 1, 1,
	2, 1, 2, 2, 1,
}

func TestA000002Name(t *testing.T) {
	a := &A000002{}

	name := a.Name()
	if name != "A000002" {
		t.Error("name should be A000002, got " + name)
	}
}

func TestA000002Next(t *testing.T) {
	a := &A000002{}

	for i := 0; i < 10; i++ {
		x, ok := a.Next()
		if !ok {
			t.Error("not ok at " + strconv.Itoa(i))
		}

		if x != testA000002[i] {
			t.Error("x should be " + strconv.Itoa(testA000002[i]) + ", got " + strconv.Itoa(x))
		}
	}
}

func TestA000002Reset(t *testing.T) {
	a := &A000002{}

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

	if x != testA000002[0] {
		t.Error("x should be " + strconv.Itoa(testA000002[0]) + ", got " + strconv.Itoa(x))
	}
}
