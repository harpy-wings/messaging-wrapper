package ampqc

import (
	w "github.com/harpy-wings/messaging-wrapper"
	amqp "github.com/rabbitmq/amqp091-go"
)

type queue struct {
	messenger *messenger

	name string

	ch *amqp.Channel

	Config struct {
		Args       map[string]interface{}
		NoWait     bool
		Durable    bool
		AutoDelete bool
		Exclusive  bool
	}
}

var _ w.Queue = &queue{}

// QueueBind binds an exchange to a queue so that publishings to the exchange will
// be routed to the queue when the publishing routing key matches the binding
// routing key.
func (s *queue) Bind(ex w.Exchange, key string, ops ...w.BindOption) error {
	return s.ch.QueueBind(s.name, key, ex.Name(), s.Config.NoWait, s.Config.Args)
}

// Unbind removes a binding between an exchange and queue matching the key and
// arguments.
func (s *queue) Unbind(ex w.Exchange, key string) error {
	return s.ch.QueueUnbind(s.name, key, ex.Name(), s.Config.Args)
}

//	QueueDelete removes the queue from the server including all bindings then
//
// purges the messages based on server configuration, returning the number of
// messages purged.
func (s *queue) Delete(ifEmpty, ifUnused bool) (int, error) {
	n, err := s.ch.QueueDelete(s.name, ifUnused, ifEmpty, false)
	if err != nil {
		return n, err
	}
	delete(s.messenger.queues, s.name)
	return n, nil
}

//	QueuePurge removes all messages from the named queue which are not waiting to
//
// be acknowledged.  Messages that have been delivered but have not yet been
// acknowledged will not be removed.
func (s *queue) Purge() (int, error) {
	return s.ch.QueuePurge(s.name, false)
}

// number of consumers receiving deliveries
func (s *queue) Consumers() (int, error) {
	q, err := s.ch.QueueInspect(s.name)
	if err != nil {
		return 0, err
	}
	return q.Consumers, nil
}

// count of messages not awaiting acknowledgment
func (s *queue) Messages() (int, error) {
	q, err := s.ch.QueueInspect(s.name)
	if err != nil {
		return 0, err
	}
	return q.Messages, nil
}

// server confirmed or generated name
func (s *queue) Name() string {
	return s.name
}

func (s *queue) loadDefault() {
	// todo
}
