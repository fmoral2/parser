package rabbit

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func Consumer() {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatal(err)
	}

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

	file, err := os.OpenFile("./input/message.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	forever := make(chan bool)
	fmt.Println("Waiting messages")

	go func() {
		for msg := range msgs {
			fmt.Printf("Received Message: %s\n", msg.Body)
			_, _ = file.Write(msg.Body)
		}
	}()

	JsonToCsv()

	<-forever
}
