package concurrent

import "fmt"

const introChannelClose = `
Closing a channel indicates that no more values will be sent on it.
This can be useful to communicate completion to the channel’s receivers.`

func playChannelClose() {
	fmt.Println("===channel close===")

	fmt.Println(introChannelClose)

	q := make(chan int, 1)
	done := make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			q <- i
		}
		close(q)
	}()

	go func() {
		for {
			v, ok := <-q

			if ok {
				fmt.Println("value v=", v)
			} else {
				fmt.Println("close v=", v)
				done <- true
				return
			}
		}
	}()

	<-done
}
