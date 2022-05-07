package function

import "fmt"

func play_recursion() {
	fmt.Println("===recursion===")

	fmt.Println("fib 0=", fib(0))
	fmt.Println("fib 1=", fib(1))
	fmt.Println("fib 2=", fib(2))
	fmt.Println("fib 3=", fib(3))
	fmt.Println("fib 4=", fib(4))

}

func fib(n int) int {
	if n < 2 {
		return 1
	}

	return fib(n-2) + fib(n-1)
}
