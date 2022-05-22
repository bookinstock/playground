package concurrent

import "fmt"

const introChannelDirection = `
When using channels as function parameters,
you can specify if a channel is meant to only send or receive values.
This specificity increases the type-safety of the program.`

func playChannelDirection() {
	fmt.Println("===channel direction===")
	fmt.Println(introChannelDirection)

	q1 := make(chan int)
	q2 := make(chan int)

	go func(in <-chan int, out chan<- int) {
		q2 <- (<-q1) + 1
	}(q1, q2)

	q1 <- 1

	r := <-q2

	fmt.Println("r=", r)

}
