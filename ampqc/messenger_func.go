package ampqc

import (
	"errors"

	w "github.com/harpy-wings/messaging-wrapper"
)

func (s *messenger) leadDefault() {
	s.config.autoAck = defaultAutoAck
	s.config.exclusive = defaultExclusive
	s.config.noLocal = defaultNoLocal
}

func (s *messenger) init() error { return errors.New("unimplemented") }

// NewConsumer creates a new consumer for this messenger.
func (s *messenger) NewConsumer(ops ...w.ConsumerOption) (w.Consumer, error) {
	c := new(consumer)
	var err error
	c.loadDefault()
	for _, fn := range ops {
		err = fn(c)
		if err != nil {
			return nil, err
		}
	}

	//err = c.init()
	c.name = c.Config.Name

	c.ch, err = s.conn.Channel()
	if err != nil {
		return nil, err
	}

	c.c, err = c.ch.Consume(c.Config.Queue.Name(), c.name, false, c.Config.Exclusive, c.Config.NoLocal, c.Config.NoWait, c.Config.Args)
	if err != nil {
		return nil, err
	}
	c.messenger = s 
	
	s.consumers[c.name] = c

	return c, nil
}

func (s *messenger) GetConsumer(name string) (w.Consumer, error)

// Producer create a new producer for this messenger.
func (s *messenger) NewProducer(ops ...w.ProducerOption) (w.Producer, error)
func (s *messenger) GetProducer(name string) (w.Producer, error)

// NewExchange create a new Exchange for this messenger.
func (s *messenger) NewExchange(ops ...w.ExchangeOption) (w.Exchange, error)
func (s *messenger) GetExchange(name string) (w.Exchange, error)

// NewQueue create a new Queue for this messenger.
func (s *messenger) NewQueue(ops ...w.QueueOption) (w.Queue, error) {
	q := new(queue)
	var err error
	q.loadDefault()
	for _, fn := range ops {
		err = fn(q)
		if err != nil {
			return nil, err
		}
	}
	// err = q.init()

	q.ch, err = s.conn.Channel()
	if err != nil {
		return nil, err
	}
	Q, err := q.ch.QueueDeclare(q.name, q.Config.Durable, q.Config.AutoDelete, q.Config.Exclusive, q.Config.NoWait, q.Config.Args)
	if err != nil {
		return nil, err
	}
	q.name = Q.Name

	s.queues[q.name] = q

	return q, nil
}

func (s *messenger) GetQueue(name string) (w.Queue, error) {
	q, ok := s.queues[name]
	if !ok {
		return nil, w.ErrorNotFound("queue[" + name + "]")
	}
	return q, nil
}

// CancelConsumers stops deliveries to all consumers established.
func (s *messenger) CancelConsumers() error {
	for _, v := range s.consumers {
		// todo
		_ = v
	}
	return nil
}

// GracefulShutdown shutdowns the Messenger.
func (s *messenger) GracefulShutdown() error

// Wait locks the caller routine until the Messenger gracefully shutdown.
func (s *messenger) Wait() {
	foreverTmp := make(chan struct{})
	<-foreverTmp
}
