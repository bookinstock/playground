package concurrent

import (
	"fmt"
	"time"
)

const introChannelSelect = "select lets you wait on multiple channel operations"

func playChannelSelect() {
	fmt.Println("===channel select===")

	fmt.Println(introChannelSelect)

	q1 := make(chan int)
	q2 := make(chan int)

	go func(q chan<- int) {
		time.Sleep(time.Millisecond)
		q <- 1
	}(q1)

	go func(q chan<- int) {
		time.Sleep(time.Millisecond * 2)
		q <- 2
	}(q2)

	for i := 0; i < 2; i++ {
		select {
		case msg := <-q1:
			fmt.Println("case 1 msg=", msg)
		case msg := <-q2:
			fmt.Println("case 2 msg=", msg)
		}
	}

}
