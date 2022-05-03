package main

import (
	"fmt"
	"testapp/control_flow"
	"testapp/data_type"
)

func main() {
	foo()

	data_type.Run()

	control_flow.Run()
}

func foo() {
	fmt.Println("foo")
}
