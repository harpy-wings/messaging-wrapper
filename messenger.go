package messagingwrapper

// Messenger
// TODO - add proper comments.
type Messenger interface {

	// TODO - add proper comments.
	NewConsumer(ops ...ConsumerOption) (Consumer, error)

	// TODO - add proper comments.
	NewProducer() (Producer, error)

	// TODO - add proper comments.
	NewExchange(ops ...ExchangeOption) (Exchange, error)

	// TODO - add proper comments.
	NewQueue(ops ...QueueOption) (Queue, error)

	// GracefulShutdown shutdowns the Messenger.
	GracefulShutdown() error

	// Wait locks the caller routine until the Messenger gracefully shutdown.
	Wait()
}
