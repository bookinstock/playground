package cmd

import (
	"github.com/urfave/cli"
)

var Commands = []cli.Command{
	{
		Name:  "simple",
		Usage: "mode simple",
		Action: func(c *cli.Context) error {

			return nil
		},
	},
	{
		Name:  "worker",
		Usage: "mode worker",
		Action: func(c *cli.Context) error {
			return nil
		},
	},
}
