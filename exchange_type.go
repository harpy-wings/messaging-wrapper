package messagingwrapper

type ExchangeType int

const (
	ExchangeType_UNKNOWN = iota
	ExchangeType_DIRECT
	ExchangeType_FANOUT
	ExchangeType_TOPIC
	ExchangeType_HEADERS
)

var (
	_ExchangeType2String = map[ExchangeType]string{
		ExchangeType_UNKNOWN: "unknown",
		ExchangeType_DIRECT:  "direct",
		ExchangeType_FANOUT:  "fanout",
		ExchangeType_TOPIC:   "topic",
		ExchangeType_HEADERS: "headers",
	}

	_String2ExchangeType = map[string]ExchangeType{
		"direct":  ExchangeType_DIRECT,
		"fanout":  ExchangeType_FANOUT,
		"topic":   ExchangeType_TOPIC,
		"headers": ExchangeType_HEADERS,
	}
)

func (e ExchangeType) String() string {
	s, ok := _ExchangeType2String[e]
	if !ok {
		return "unknown"
	}
	return s
}

//ExchangeTypeFromString gets the ExchangeType value from a string
//
// "direct", "fanout", "topic" and "headers".
func (e ExchangeType) ExchangeTypeFromString(s string) ExchangeType {
	e, ok := _String2ExchangeType[s]
	if !ok {
		return ExchangeType_UNKNOWN
	}
	return e
}
