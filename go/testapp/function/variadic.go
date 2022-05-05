package function

import "fmt"

func variadic(lst ...int) {
	for i := 0; i < len(lst); i++ {
		fmt.Print(lst[i], " ")
	}
	fmt.Println()
}

func play_variadic() {
	variadic([]int{1, 2, 3, 4, 5}...)
}
