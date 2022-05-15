package cmd

import (
	"rabbit/role"

	"github.com/urfave/cli"
)

var Flags = []cli.Flag{
	cli.StringFlag{
		Name:        "r",
		Value:       role.ALL,
		Usage:       "role",
		Destination: &role.Role,
	},
}
