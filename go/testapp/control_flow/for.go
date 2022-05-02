package control_flow

import "fmt"

func play_for() {
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	i := 0
	for i < 5 {
		fmt.Print(i, " ")
		i += 1
	}
	fmt.Println()

}
