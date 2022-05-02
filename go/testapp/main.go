package main

import (
	"fmt"
	"testapp/go/testapp/control_flow"
	"testapp/go/testapp/data_type"
)

func main() {
	foo()

	data_type.Run()

	control_flow.Run()
}

func foo() {
	fmt.Println("foo")
}
