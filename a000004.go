package goeis

// The zero sequence.
type A000004 struct {}

func (a *A000004) Name() string {
	return "A000004"
}

func (a *A000004) Next() (n int, ok bool) {
	ok = true
	return
}

func (a *A000004) Reset() {}
