package data_type

import "fmt"

func play_slice() {
	s := make([]int, 3)
	fmt.Println("s=", s)

	s[0] = 1
	s[1] = 2
	s[2] = 3
	fmt.Println("s=", s)

	fmt.Println("len=", len(s))

	s = append(s, 4)
	s = append(s, 5, 6)
	s = append(s, []int{7, 8, 9}...)
	fmt.Println("s=", s)
}
