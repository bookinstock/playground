package control_flow

import "fmt"

func play_if() {
	a := 1

	if a < 1 {
		fmt.Println("a < 1")
	} else if a == 1 {
		fmt.Println("a == 1")
	} else {
		fmt.Println("a > 1")
	}
}
