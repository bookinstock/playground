package function

import "fmt"

func play_variadic() {
	sum := variadic_add([]int{1, 2, 3, 4, 5}...)
	fmt.Println("sum=", sum)
}

func variadic_add(lst ...int) int {
	sum := 0
	for _, e := range lst {
		sum += e
	}
	return sum
}
