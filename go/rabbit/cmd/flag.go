package cmd

import "github.com/urfave/cli"

var ConfigPath string

var Flags = []cli.Flag{
	cli.StringFlag{
		Name:        "c",
		Value:       "test",
		Usage:       "config path",
		Destination: &ConfigPath,
	},
}
