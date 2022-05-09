package object

import "fmt"

func play_embed() {
	fmt.Println("===embed===")
	a := A{Base: Base{V: 1, II: F1{f: 2}}, VA: 3}
	b := B{Base: Base{V: 1, II: F2{f: 4}}, VB: 5}
	fmt.Printf("a=%#v\n", a)
	fmt.Printf("b=%#v\n", b)
	fmt.Println("a v=", a.v())
	fmt.Println("b v=", b.v())
}

type Base struct {
	V int

	II
}

func (base Base) value() int {
	return base.V
}

type A struct {
	Base

	VA int
}

type B struct {
	Base

	VB int
}

func (a *A) v() int {
	return a.value() + a.VA + a.foo()
}

func (b *B) v() int {
	return b.value() + b.VB + b.foo()
}

type II interface {
	foo() int
}

type F1 struct {
	f int
}

type F2 struct {
	f int
}

func (f F1) foo() int {
	return f.f
}

func (f F2) foo() int {
	return f.f
}
