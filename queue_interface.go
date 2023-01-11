package messagingwrapper

// Queues in the AMQP 0-9-1 model are very similar to queues in other message-
// and task-queueing systems: they store messages that are consumed by applications.
type Queue interface {

	// QueueBind binds an exchange to a queue so that publishings to the exchange will
	// be routed to the queue when the publishing routing key matches the binding
	// routing key.
	Bind(ex Exchange, key string, ops ...BindOption) error

	// Unbind removes a binding between an exchange and queue matching the key and
	// arguments.
	Unbind(ex Exchange, key string) error

	// 	QueueDelete removes the queue from the server including all bindings then
	// purges the messages based on server configuration, returning the number of
	// messages purged.
	Delete(ifEmpty, ifUnused bool) (int, error)

	// 	QueuePurge removes all messages from the named queue which are not waiting to
	// be acknowledged.  Messages that have been delivered but have not yet been
	// acknowledged will not be removed.
	Purge() (int, error)

	// number of consumers receiving deliveries
	Consumers() (int, error)

	// count of messages not awaiting acknowledgment
	Messages() (int, error)

	// server confirmed or generated name
	Name() string
}
