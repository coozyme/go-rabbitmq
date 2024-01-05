package main

import (
	"log"

	"github.com/streadway/amqp"
)

func errorWrapper(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	errorWrapper(err, "Failed to connect rabbitmq")
	defer conn.Close()

	ch, err := conn.Channel()
	errorWrapper(err, "Failed to open a channel")
	defer ch.Close()

	errorWrapper(err, "Failed to declare a queue")

	body := "Hi halovina, keep in touch"
	err = ch.Publish(
		"exch.testing.topic", // exchange
		"DELL_TESTING",       // routing key
		false,                // mandatory
		false,                // immadiate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	errorWrapper(err, "Failed to publish message")
	log.Printf("Sending message success: %s", body)

}
