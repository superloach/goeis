package goeis

import (
	"strconv"
	"testing"
)

var testA000001 = []int{
	0, 1, 1, 1, 2,
	1, 2, 1, 5, 2,
}

func TestA000001Name(t *testing.T) {
	a := &A000001{}

	name := a.Name()
	if name != "A000001" {
		t.Error("name should be A000001, got " + name)
	}
}

func TestA000001Next(t *testing.T) {
	a := &A000001{}

	for i := 0; i < 10; i++ {
		x, ok := a.Next()
		if !ok {
			t.Error("not ok at " + strconv.Itoa(i))
		}

		if x != testA000001[i] {
			t.Error("x should be " + strconv.Itoa(testA000001[i]) + ", got " + strconv.Itoa(x))
		}
	}
}

func TestA000001Reset(t *testing.T) {
	a := &A000001{}

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

	if x != testA000001[0] {
		t.Error("x should be " + strconv.Itoa(testA000001[0]) + ", got " + strconv.Itoa(x))
	}
}
