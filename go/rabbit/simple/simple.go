package simple

import (
	"fmt"
	"rabbit/cmd"

	"golang.org/x/sync/errgroup"
)

func Run() {

	if cmd.IsProducer() {
		send()
	} else if cmd.IsConsumer() {
		receive()
	} else {
		g := new(errgroup.Group)

		g.Go(func() error {
			send()
			return nil
		})

		g.Go(func() error {
			receive()
			return nil
		})

		if err := g.Wait(); err == nil {
			fmt.Println("success")
		} else {
			fmt.Println("fail e=", err)
		}
	}
}
