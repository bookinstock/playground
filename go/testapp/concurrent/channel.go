package concurrent

import "fmt"

const introChannel = `
Channels are the pipes that connect concurrent goroutines.
You can send values into channels from one goroutine
and receive those values into another goroutine.`

func playChannel() {
	fmt.Println("===channel===")

	fmt.Println(introChannel)

	queue := make(chan string)

	go func() {
		queue <- "ping"
	}()

	msg := <-queue

	fmt.Println("msg=", msg)
}
