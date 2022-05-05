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

	i = 0
	for {
		if i == 2 {
			i += 1
			continue
		}
		if i == 5 {
			break
		}
		fmt.Print(i, " ")
		i += 1
	}
	fmt.Println()

	lst := []int{1, 2, 3, 4, 5}
	for i, e := range lst {
		fmt.Println("i=", i, "e=", e)
	}
	fmt.Println()

	play_range()
}

func play_range() {
	fmt.Println("===range===")

	s := []int{1, 2, 3}
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	str := "foobar"

	for i, e := range s {
		fmt.Println("i=", i, "e=", e)
	}
	fmt.Println()

	for k, v := range m {
		fmt.Println("k=", k, "v=", v)
	}
	fmt.Println()

	for i, e := range str {
		fmt.Println("i=", i, "e=", e)
	}
	fmt.Println()
}
