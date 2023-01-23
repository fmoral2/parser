package rabbit

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func Consumer() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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

	file, errs := os.OpenFile("input/message.json", os.O_RDWR|os.O_CREATE, 0755)
	if errs != nil {
		fmt.Println(errs)
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
