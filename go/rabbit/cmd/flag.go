package cmd

import "github.com/urfave/cli"

const (
	PRODUCER = "p"
	CONSUMER = "c"
)

var role string

var Flags = []cli.Flag{
	cli.StringFlag{
		Name:        "r",
		Value:       CONSUMER,
		Usage:       "role",
		Destination: &role,
	},
}

func IsProducer() bool {
	return role == PRODUCER
}

func IsConsumer() bool {
	return role == CONSUMER
}
