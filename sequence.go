package goeis

type Sequence interface {
	Name() string
	Reset()
	Next() (int, bool)
}

var Sequences = map[int]Sequence{
	1: &A000001{},
	2: &A000002{},
	3: &A000003{},
	4: &A000004{},
	5: &A000005{},
	6: &A000006{},
	40: &A000040{},
}
