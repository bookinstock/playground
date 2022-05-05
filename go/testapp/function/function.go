package function

import "fmt"

func Run() {

	fmt.Println("===func===")

	fmt.Println("add=", add(1, 2))

	a, b := add_sub(1, 2)
	fmt.Println("add_sub=", a, b)

	play_variadic()
}

func add(a, b int) int {
	return a + b
}

func add_sub(a, b int) (int, int) {
	return a + b, a - b
}
