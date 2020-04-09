package goeis

import (
	"strconv"
	"testing"
)

var testA000004 = []int{
	0, 0, 0, 0, 0,
	0, 0, 0, 0, 0,
}

func TestA000004Name(t *testing.T) {
	a := &A000004{}

	name := a.Name()
	if name != "A000004" {
		t.Error("name should be A000004, got " + name)
	}
}

func TestA000004Next(t *testing.T) {
	a := &A000004{}

	for i := 0; i < 10; i++ {
		x, ok := a.Next()
		if !ok {
			t.Error("not ok at " + strconv.Itoa(i))
		}

		if x != testA000004[i] {
			t.Error("x should be " + strconv.Itoa(testA000004[i]) + ", got " + strconv.Itoa(x))
		}
	}
}

func TestA000004Reset(t *testing.T) {
	a := &A000004{}

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

	if x != testA000004[0] {
		t.Error("x should be " + strconv.Itoa(testA000004[0]) + ", got " + strconv.Itoa(x))
	}
}
