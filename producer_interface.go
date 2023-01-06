package messagingwrapper

import "context"

type Producer interface {
	C() chan Message

	Publish(context.Context, Message)
}
