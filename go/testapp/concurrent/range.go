package concurrent

import "fmt"

func playRange() {
	fmt.Println("===range===")

	q := make(chan int, 1)

	go func() {
		i := 0
		for i < 5 {
			q <- i
			i++
		}
		close(q)
	}()

	for e := range q {
		fmt.Println("value e=", e)
	}

}
