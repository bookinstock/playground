package main

import (
	"fmt"
	"testapp/go/testapp/data_type"
)

func main() {
	foo()

	data_type.Run()
}

func foo() {
	fmt.Println("foo")
}
