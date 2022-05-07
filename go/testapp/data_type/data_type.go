package data_type

import "fmt"

const (
	const_a = 1
	const_b = 2
)

var (
	var_a = 1
	var_b = 2
)

func Run() {
	checkVarConst()

	playBool()
	playNumber()
	playString()
	play_slice()
	play_map()

	play_pointer()
}

func checkVarConst() {
	fmt.Println("const=", const_a, const_b)
	fmt.Println("var=", var_a, var_b)
}
