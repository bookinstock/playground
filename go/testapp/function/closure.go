package function

import "fmt"

func play_closure() {

	fmt.Println("===closure===")

	c := gen_counter()

	fmt.Println("c = ", c())
	fmt.Println("c = ", c())
	fmt.Println("c = ", c())

}

func gen_counter() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
