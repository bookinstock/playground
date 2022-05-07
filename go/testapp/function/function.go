package function

import "fmt"

func Run() {

	fmt.Println("===func===")

	fmt.Println("add=", add(1, 2))

	a, b := add_sub(1, 2)
	fmt.Println("add_sub=", a, b)

	play_variadic()

	play_closure()

	play_calculator()

	play_recursion()
}

func add(a, b int) int {
	return a + b
}

func add_sub(a, b int) (int, int) {
	return a + b, a - b
}

func play_calculator() {
	a := 1
	b := 2

	add := func(x, y int) int {
		return x + y
	}

	sub := func(x, y int) int {
		return x - y
	}

	mul := func(x, y int) int {
		return x * y
	}

	div := func(x, y int) int {
		return x / y
	}

	fmt.Println("add =", calculator(a, b, add))
	fmt.Println("sub =", calculator(a, b, sub))
	fmt.Println("mul =", calculator(a, b, mul))
	fmt.Println("div =", calculator(a, b, div))
}

func calculator(a, b int, fn func(int, int) int) int {
	return fn(a, b)
}
