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

	c := make([]int, len(s))
	copy(c, s)
	fmt.Println("c=", c)

	fmt.Println("s-2-5=", s[2:5])
	fmt.Println("s-2=", s[2:])
	fmt.Println("s-5=", s[:5])
	fmt.Println("s-=", s[:])

	v := 1
	ss := make([][]int, 3)
	for i := range ss {
		ss[i] = make([]int, 3)
		for j := range ss[i] {
			ss[i][j] = v
			v += 1
		}
	}
	fmt.Println("ss=", ss)
}

// https://go.dev/blog/slices-intro
