package concurrent

import (
	"fmt"
	"time"
)

const introNonBlockingChannel = `
Basic sends and receives on channels are blocking.
use select with a default clause to implement
non-blocking sends, receives, and even non-blocking multi-way selects.
`

func playChannelNonBlocking() {
	fmt.Println("===non-blocking channel===")
	fmt.Println(introNonBlockingChannel)

	// sender
	q := make(chan int)
	go func() {
		q <- 1
	}()
	select {
	case <-q:
		fmt.Println("case")
	default:
		fmt.Println("default")
	}

	// main wait
	q = make(chan int)
	go func() {
		q <- 1
	}()
	time.Sleep(time.Millisecond)
	select {
	case <-q:
		fmt.Println("case")
	default:
		fmt.Println("default")
	}

	// sender wait
	q = make(chan int)
	go func() {
		time.Sleep(time.Millisecond)
		q <- 1
	}()
	select {
	case <-q:
		fmt.Println("case")
	default:
		fmt.Println("default")
	}

	// no sender
	q = make(chan int, 1)
	select {
	case <-q:
		fmt.Println("case")
	default:
		fmt.Println("default")
	}

	// no sender buffer
	q = make(chan int, 1)
	select {
	case <-q:
		fmt.Println("case")
	default:
		fmt.Println("default")
	}

	// receiver
	q = make(chan int)
	go func() {
		<-q
	}()
	select {
	case q <- 1:
		fmt.Println("case")
	default:
		fmt.Println("default")
	}

	// main wait
	q = make(chan int)
	go func() {
		<-q
	}()
	time.Sleep(time.Millisecond)
	select {
	case q <- 1:
		fmt.Println("case")
	default:
		fmt.Println("default")
	}

	// no receiver
	q = make(chan int)
	select {
	case q <- 1:
		fmt.Println("case")
	default:
		fmt.Println("default")
	}

	// no receiver buffer
	q = make(chan int, 1)
	select {
	case q <- 1:
		fmt.Println("case")
	default:
		fmt.Println("default")
	}

	// multi sender and receiver
	in := make(chan int, 1)
	out := make(chan int, 1)
	select {
	case in <- 1:
		fmt.Println("case in")
	case <-out:
		fmt.Println("case out")
	default:
		fmt.Println("default")
	}
}
