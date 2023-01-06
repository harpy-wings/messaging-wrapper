package messagingwrapper

// Consume immediately starts delivering queued messages.
type Consumer interface {
	// Foreach executes a function fot the message and waits the process to be finished,
	// if error be nil then call the Ack. of the message
	Foreach(func(msg Message) error)

	// ForeachN is a bounded parallelism mechanism to process the message in N number of goroutine.
	// if error be nil then call the Ack. of the message
	ForeachN(fn func(msg Message) error, n int)
}
