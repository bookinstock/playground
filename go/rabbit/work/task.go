package work

import (
	"fmt"
	"rabbit/utils"
	"time"

	"github.com/streadway/amqp"
)

func publish() {
	fmt.Println("publish")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	utils.FailOnError(err, "Failed to declare a queue")

	i := 1
	for {
		msg := fmt.Sprintf("foo-%d", i)
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         []byte(msg),
			})
		utils.FailOnError(err, "Failed to publish a message")

		// fmt.Printf("publish msg=%s\n", msg)

		time.Sleep(time.Millisecond * 200)
		i += 1
	}
}
