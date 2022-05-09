package object

import "fmt"

func play_interface() {
	t1 := T1{1}
	t2 := T1{2}
	fmt.Println("t1", getValue(t1))
	fmt.Println("t2", getValue(t2))

	var t T = t1
	fmt.Println("t=", getValue(t))
}

type T1 struct {
	v int
}

type T2 struct {
	v int
}

func (t T1) value() int {
	return t.v
}

func (t T2) value() int {
	return t.v
}

type T interface {
	value() int
}

func getValue(t T) int {
	return t.value()
}
