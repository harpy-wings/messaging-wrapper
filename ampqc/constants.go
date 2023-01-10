package ampqc

const (
	defaultAutoAck bool = false

	// When exclusive is true, the server will ensure that this is the sole consumer
	// from this queue. When exclusive is false, the server will fairly distribute
	// deliveries across multiple consumers.
	defaultExclusive bool = false

	defaultNoLocal bool = false // is not supported

	defaultNoWait bool = false
)
