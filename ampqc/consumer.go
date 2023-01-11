package ampqc

import (
	"log"

	w "github.com/harpy-wings/messaging-wrapper"
	amqp "github.com/rabbitmq/amqp091-go"
)

type consumer struct {
	messenger *messenger

	name string

	ch *amqp.Channel

	c <-chan amqp.Delivery

	Config struct {
		Args      map[string]interface{}
		NoWait    bool
		NoLocal   bool
		Exclusive bool
		Name      string
		Queue     w.Queue
	}
}

// Foreach executes a function fot the message and waits the process to be finished,
// if error be nil then call the Ack. of the message
func (s *consumer) Foreach(fn func(msg w.Message) error) {
	s.ForeachN(fn, 1)
}

// ForeachN is a bounded parallelism mechanism to process the message in N number of goroutine.
// if error be nil then call the Ack. of the message
func (s *consumer) ForeachN(fn func(msg w.Message) error, n int) {
	for i := 0; i < n; i++ {
		go func() {
			for m := range s.c {
				err := fn(m.Body)
				if err == nil {
					m.Ack(true)
				}
				if err != nil {
					log.Println(err)
				}
			}
		}()
	}
}

func (s *consumer) loadDefault() {
	//todo
}
