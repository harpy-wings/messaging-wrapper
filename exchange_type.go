package messagingwrapper

type ExchangeType int

// "direct", "fanout", "topic" and "headers".
const (
	ExchangeType_UNKNOWN = iota
	ExchangeType_DIRECT
	ExchangeType_FANOUT
	ExchangeType_TOPIC
	ExchangeType_HEADERS
)
