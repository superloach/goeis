package goeis

import (
	"strconv"
	"testing"
)

var testA000007 = []int{
	1, 0, 0, 0, 0,
	0, 0, 0, 0, 0,
}

func TestA000007Name(t *testing.T) {
	a := &A000007{}

	name := a.Name()
	if name != "A000007" {
		t.Error("name should be A000007, got " + name)
	}
}

func TestA000007Next(t *testing.T) {
	a := &A000007{}

	for i := 0; i < 10; i++ {
		x, ok := a.Next()
		if !ok {
			t.Error("not ok at " + strconv.Itoa(i))
		}

		if x != testA000007[i] {
			t.Error("x should be " + strconv.Itoa(testA000007[i]) + ", got " + strconv.Itoa(x))
		}
	}
}

func TestA000007Reset(t *testing.T) {
	a := &A000007{}

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

	if x != testA000007[0] {
		t.Error("x should be " + strconv.Itoa(testA000007[0]) + ", got " + strconv.Itoa(x))
	}
}
