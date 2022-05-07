package object

func play_embed() {

}

type A struct {
	v int
}

func (a A) value() int {
	return a.v
}

type AA struct {
	A
}
