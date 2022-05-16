package concurrent

import (
	"fmt"
	"time"
)

const introTimeout = `
Timeouts are important for programs that connect to external resources
or that otherwise need to bound execution time.
`

func playTimeout() {
	fmt.Println("===timeout===")
	fmt.Println(introTimeout)

	q := make(chan int)

	go func() {
		time.Sleep(time.Millisecond)
		q <- 1
	}()
	select {
	case <-q:
		fmt.Println("get q data")
	case <-time.After(time.Millisecond * 2):
		fmt.Println("timeout")
	}

	go func() {
		time.Sleep(time.Millisecond * 2)
		q <- 1
	}()
	select {
	case <-q:
		fmt.Println("get q data")
	case <-time.After(time.Millisecond):
		fmt.Println("timeout")
	}

}
