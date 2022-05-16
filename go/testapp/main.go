package main

import (
	"fmt"
	"testapp/concurrent"
	"testapp/control_flow"
	"testapp/data_type"
	"testapp/function"
	"testapp/object"
)

func main() {
	foo()

	data_type.Run()

	control_flow.Run()

	function.Run()

	// err.Run()

	object.Run()

	concurrent.Run()
}

func foo() {
	fmt.Println("foo")
}
