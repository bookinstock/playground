package work

import (
	"fmt"
	"rabbit/role"

	"golang.org/x/sync/errgroup"
)

func Run() {

	if role.IsProducer() {
		publish()
	} else if role.IsConsumer() {
		consume()
	} else {
		g := new(errgroup.Group)

		g.Go(func() error {
			publish()
			return nil
		})

		g.Go(func() error {
			consume()
			return nil
		})

		if err := g.Wait(); err == nil {
			fmt.Println("success")
		} else {
			fmt.Println("fail e=", err)
		}
	}
}
