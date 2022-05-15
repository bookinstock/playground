package simple

import (
	"fmt"
	"rabbit/role"

	"golang.org/x/sync/errgroup"
)

func Run() {

	if role.IsProducer() {
		send()
	} else if role.IsConsumer() {
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
