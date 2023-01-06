package messagingwrapper

type BindOption func() error

type ConsumerOption func() error

type ExchangeOption func() error

type QueueOption func() error

type ProducerOption func() error

func ProducerOptionOnError(func(Message)) ProducerOption
func ProducerOptionWithRetry(n int) ProducerOption
func ProducerOptionWithExchange(Exchange) Producer
