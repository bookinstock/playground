package main

import (
	"fmt"
	"os"
	"rabbit/cmd"

	"github.com/urfave/cli"
)

// func main() {
// 	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
// 	failOnError(err, "Failed to connect to RabbitMQ")
// 	defer conn.Close()
// }

// func failOnError(err error, msg string) {
// 	if err != nil {
// 		log.Fatalf("%s: %s", msg, err)
// 	}
// }

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
