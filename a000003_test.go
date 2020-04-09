package goeis

import (
	"strconv"
	"testing"
)

var testA000003 = []int{
	1, 1, 1, 1, 2,
	2, 1, 2, 2, 2,
}

func TestA000003Name(t *testing.T) {
	a := &A000003{}

	name := a.Name()
	if name != "A000003" {
		t.Error("name should be A000003, got " + name)
	}
}

func TestA000003Next(t *testing.T) {
	a := &A000003{}

	for i := 0; i < 10; i++ {
		x, ok := a.Next()
		if !ok {
			t.Error("not ok at " + strconv.Itoa(i))
		}

		if x != testA000003[i] {
			t.Error("x should be " + strconv.Itoa(testA000003[i]) + ", got " + strconv.Itoa(x))
		}
	}
}

func TestA000003Reset(t *testing.T) {
	a := &A000003{}

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

	if x != testA000003[0] {
		t.Error("x should be " + strconv.Itoa(testA000003[0]) + ", got " + strconv.Itoa(x))
	}
}
