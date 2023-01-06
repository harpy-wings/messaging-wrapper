package messagingwrapper

type BindOption func() error

func BindOptionWithArgs(map[string]interface{}) BindOption
func BindOptionWithNoWait(b bool) BindOption

type ConsumerOption func() error

func ConsumerOptionWithQueue(Queue) ConsumerOption
func ConsumerOptionWithName(s string) ConsumerOption
func ConsumerOptionWithExclusive(b bool) ConsumerOption
func ConsumerOptionWithArgs(map[string]interface{}) ConsumerOption

type ExchangeOption func() error

func ExchangeOptionWithName(s string) ExchangeOption
func ExchangeOptionWithType(t ExchangeType) ExchangeOption

// Durable and Non-Auto-Deleted exchanges will survive server restarts and remain
// declared when there are no remaining bindings.  This is the best lifetime for
// long-lived exchange configurations like stable routes and default exchanges.
//
// Non-Durable and Auto-Deleted exchanges will be deleted when there are no
// remaining bindings and not restored on server restart.  This lifetime is
// useful for temporary topologies that should not pollute the virtual host on
// failure or after the consumers have completed.
//
// Non-Durable and Non-Auto-deleted exchanges will remain as long as the server is
// running including when there are no remaining bindings.  This is useful for
// temporary topologies that may have long delays between bindings.
//
// Durable and Auto-Deleted exchanges will survive server restarts and will be
// removed before and after server restarts when there are no remaining bindings.
// These exchanges are useful for robust temporary topologies or when you require
// binding durable queues to auto-deleted exchanges.
func ExchangeOptionAutoDelete(b bool) ExchangeOption
func ExchangeOptionDurable(b bool) ExchangeOption

// Exchanges declared as `internal` do not accept accept publishings. Internal
// exchanges are useful when you wish to implement inter-exchange topologies
// that should not be exposed to users of the broker.
func ExchangeOptionInternal(b bool) ExchangeOption

// A passive exchange is assumed by RabbitMQ to already exist, and attempting to connect to a
// non-existent exchange will cause RabbitMQ to throw an exception. This function
// can be used to detect the existence of an exchange.
func ExchangeOptionPassive(b bool) ExchangeOption

type QueueOption func() error

type ProducerOption func() error

func ProducerOptionOnError(func(Message)) ProducerOption
func ProducerOptionWithRetry(n int) ProducerOption
func ProducerOptionWithExchange(Exchange) Producer
