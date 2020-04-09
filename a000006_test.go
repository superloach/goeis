package goeis

import (
	"strconv"
	"testing"
)

var testA000006 = []int{
	1, 1, 2, 2, 3,
	3, 4, 4, 4, 5,
}

func TestA000006Name(t *testing.T) {
	a := &A000006{}

	name := a.Name()
	if name != "A000006" {
		t.Error("name should be A000006, got " + name)
	}
}

func TestA000006Next(t *testing.T) {
	a := &A000006{}

	for i := 0; i < 10; i++ {
		x, ok := a.Next()
		if !ok {
			t.Error("not ok at " + strconv.Itoa(i))
		}

		if x != testA000006[i] {
			t.Error("x should be " + strconv.Itoa(testA000006[i]) + ", got " + strconv.Itoa(x))
		}
	}
}

func TestA000006Reset(t *testing.T) {
	a := &A000006{}

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

	if x != testA000006[0] {
		t.Error("x should be " + strconv.Itoa(testA000006[0]) + ", got " + strconv.Itoa(x))
	}
}
