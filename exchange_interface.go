package messagingwrapper

// Exchanges are AMQP 0-9-1 entities where messages are sent to.
// Exchanges take a message and route it into zero or more queues.
// The routing algorithm used depends on the exchange type and rules called bindings.
// AMQP 0-9-1 brokers provide four exchange types
//
// - Direct exchange	(Empty string) and amq.direct \n
//
// - Fanout exchange	amq.fanout
//
// - Topic exchange	amq.topic
//
// - Headers exchange	amq.match (and amq.headers in RabbitMQ)
//
// The default exchange is a direct exchange with no name (empty string) pre-declared by the broker.
// It has one special property that makes it very useful for simple applications: every queue that is
// created is automatically bound to it with a routing key which is the same as the queue name.
type Exchange interface {

	// Bind binds this exchange to another exchange to create inter-exchange
	// routing topologies on the server.  This can decouple the private topology and
	// routing exchanges from exchanges intended solely for publishing endpoints.
	Bind(dst Exchange, key string, ops ...BindOption) error

	// Unbind unbinds the destination exchange from this exchange on the
	// server by removing the routing key between them.  This is the inverse of
	// ExchangeBind.  If the binding does not currently exist, an error will be
	// returned.
	Unbind(dst Exchange, key string)

	// Delete removes the named exchange from the server. When an exchange is
	// deleted all queue bindings on the exchange are also deleted.  If this exchange
	// does not exist, the channel will be closed with an error.
	Delete(ifUnused bool) error

	// server confirmed or generated name
	Name() string
}
