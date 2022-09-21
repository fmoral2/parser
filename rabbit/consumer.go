package rabbit

import (
	"fmt"
	"github.com/morlfm/csv_parser/application/ports"
	"github.com/streadway/amqp"
	"log"
	"os"
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

	forever := make(chan bool, 1)
	go func() {
		file, errs := os.Create("input/message.json")
		if errs != nil {
			fmt.Println(errs)
		}
		defer file.Close()
		for d := range msgs {
			fmt.Printf("Received Message: %s\n", d.Body)
			_, _ = file.Write(d.Body)
		}

	}()
	fmt.Println("Waiting messages")

	JsonToCsv()
	ports.Entry()
	<-forever

}
