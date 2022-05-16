package cmd

import (
	"rabbit/simple"
	"rabbit/work"

	"github.com/urfave/cli"
)

var Commands = []cli.Command{
	{
		Name:  "simple",
		Usage: "mode simple",
		Action: func(c *cli.Context) error {
			simple.Run()
			return nil
		},
	},
	{
		Name:  "work",
		Usage: "mode work",
		Action: func(c *cli.Context) error {
			work.Run()
			return nil
		},
	},
}
