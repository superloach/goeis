package goeis

import (
	"strconv"
	"testing"
)

var testA000040 = []int{
	2, 3, 5, 7, 11,
	13, 17, 19, 23, 29,
}

func TestA000040Name(t *testing.T) {
	a := &A000040{}

	name := a.Name()
	if name != "A000040" {
		t.Error("name should be A000040, got " + name)
	}
}

func TestA000040Next(t *testing.T) {
	a := &A000040{}

	for i := 0; i < 10; i++ {
		x, ok := a.Next()
		if !ok {
			t.Error("not ok at " + strconv.Itoa(i))
		}

		if x != testA000040[i] {
			t.Error("x should be " + strconv.Itoa(testA000040[i]) + ", got " + strconv.Itoa(x))
		}
	}
}

func TestA000040Reset(t *testing.T) {
	a := &A000040{}

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

	if x != testA000040[0] {
		t.Error("x should be " + strconv.Itoa(testA000040[0]) + ", got " + strconv.Itoa(x))
	}
}
