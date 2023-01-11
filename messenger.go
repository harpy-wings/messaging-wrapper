package messagingwrapper

// Messenger
// is the generic interface to use different message broker. each broker client must implement the Messenger interface.
type Messenger interface {

	// NewConsumer creates a new consumer for this messenger.
	NewConsumer(ops ...ConsumerOption) (Consumer, error)
	GetConsumer(name string) (Consumer, error)

	// Producer create a new producer for this messenger.
	NewProducer(ops ...ProducerOption) (Producer, error)
	GetProducer(name string) (Producer, error)

	// NewExchange create a new Exchange for this messenger.
	NewExchange(ops ...ExchangeOption) (Exchange, error)
	GetExchange(name string) (Exchange, error)

	// NewQueue create a new Queue for this messenger.
	NewQueue(ops ...QueueOption) (Queue, error)
	GetQueue(name string) (Queue, error)

	//CancelConsumers stops deliveries to all consumers established.
	CancelConsumers() error

	// GracefulShutdown shutdowns the Messenger.
	GracefulShutdown() error

	// Wait locks the caller routine until the Messenger gracefully shutdown.
	Wait()
}
