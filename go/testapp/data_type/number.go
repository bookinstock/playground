package data_type

import "fmt"

func playNumber() {
	fmt.Println("1 / 2=", 1/2)
	fmt.Println("1 / 2.0=", 1/2.0)
	fmt.Println("1 % 2=", 1%2)

	fmt.Println("1 million=", 1_000_000)
	fmt.Println("1 million=", 1e6)

	fmt.Println("int64 =", int64(100))
}
