package main

import (
	"fmt"
	"os"
	"rabbit/cmd"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "rabbit"
	app.Description = "desc"
	app.Version = "1"
	app.EnableBashCompletion = true
	app.Flags = cmd.Flags
	app.Commands = cmd.Commands

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("run err=", err)
	}

}
