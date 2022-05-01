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
	playBool()
	playNumber()
	playString()

	fmt.Println("const=", const_a, const_b)
	fmt.Println("var=", var_a, var_b)
}
