package ampqc

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type messenger struct {
	config struct {
		autoAck   bool
		exclusive bool
		noLocal   bool
		args      amqp.Table
	}
}
