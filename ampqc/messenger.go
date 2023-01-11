package ampqc

import (
	w "github.com/harpy-wings/messaging-wrapper"
	amqp "github.com/rabbitmq/amqp091-go"
)

type messenger struct {
	queues    map[string]*queue
	exchanges map[string]*exchange
	producers map[string]*producer
	consumers map[string]*consumer
	conn *amqp.Connection
	config    struct {
		autoAck   bool
		exclusive bool
		noLocal   bool
		args      map[string]interface{}
	}
}

var _ w.Messenger = &messenger{}

func New(ops ...Option) (w.Messenger, error) {
	s := new(messenger)

	s.leadDefault()

	for _, fn := range ops {
		err := fn(s)
		if err != nil {
			return nil, err
		}
	}

	err := s.init()
	if err != nil {
		return nil, err
	}

	return s, nil
}
