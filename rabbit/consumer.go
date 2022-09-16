package rabbit

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func Consumer() {
	// connect rabbit via amqp
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}

	// checkING  connection
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println("unable to consume this message")
	}

	forever := make(chan bool)
	go func() {
		file, err := os.Create("../input/message.json")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		for d := range msgs {
			fmt.Printf("Received Message: %s\n", d.Body)
			file.Write(d.Body)
		}

	}()
	fmt.Println(" [*] - Waiting for messages")

	<-forever

}
