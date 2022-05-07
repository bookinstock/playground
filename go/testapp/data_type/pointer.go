package data_type

import "fmt"

func play_pointer() {
	a := 1
	addOne(&a)
	fmt.Println("a=", a)
}

func addOne(a *int) {
	*a += 1
}
