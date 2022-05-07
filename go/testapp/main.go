package main

import (
	"fmt"
	"testapp/control_flow"
	"testapp/data_type"
	"testapp/function"
	"testapp/utils/err"
)

func main() {
	foo()

	data_type.Run()

	control_flow.Run()

	function.Run()

	err.Run()
}

func foo() {
	fmt.Println("foo")
}
