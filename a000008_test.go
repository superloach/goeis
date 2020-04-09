package goeis

import (
	"strconv"
	"testing"
)

var testA000008 = []int{
	1, 1, 2, 2, 3,
	4, 5, 6, 7, 8,
}

func TestA000008Name(t *testing.T) {
	a := &A000008{}

	name := a.Name()
	if name != "A000008" {
		t.Error("name should be A000008, got " + name)
	}
}

func TestA000008Next(t *testing.T) {
	a := &A000008{}

	for i := 0; i < 10; i++ {
		x, ok := a.Next()
		if !ok {
			t.Error("not ok at " + strconv.Itoa(i))
		}

		if x != testA000008[i] {
			t.Error("x should be " + strconv.Itoa(testA000008[i]) + ", got " + strconv.Itoa(x))
		}
	}
}

func TestA000008Reset(t *testing.T) {
	a := &A000008{}

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

	if x != testA000008[0] {
		t.Error("x should be " + strconv.Itoa(testA000008[0]) + ", got " + strconv.Itoa(x))
	}
}
