package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func t() error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	_ = q
	return nil
}

func main() {
	err := t()
	if err != nil {
		panic(err)
	}
}
