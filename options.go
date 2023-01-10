package messagingwrapper

import (
	"reflect"
)

type BindOption func(interface{}) error

func BindOptionWithArgs(v map[string]interface{}) BindOption {
	return func(vc interface{}) error {
		V := reflect.ValueOf(vc)
		if V.Kind() != reflect.Pointer {
			return ErrorNotPointer(reflect.ValueOf(vc).Kind())
		}

		CV := V.Elem().FieldByName("Config")
		if CV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String(), "Config")
		}
		VCV := CV.FieldByName("Args")
		if VCV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String()+".Config", "Args")
		}
		if VCV.Type() != reflect.TypeOf(v) {
			return ErrorMismatchType(V.Elem().Kind().String()+".Config.Args", reflect.TypeOf(v), VCV.Type())
		}
		VCV.Set(reflect.ValueOf(v))
		return nil
	}
}

func BindOptionWithNoWait(b bool) BindOption {
	return func(vc interface{}) error {
		V := reflect.ValueOf(vc)
		if V.Kind() != reflect.Pointer {
			return ErrorNotPointer(reflect.ValueOf(vc).Kind())
		}

		CV := V.Elem().FieldByName("Config")
		if CV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String(), "Config")
		}
		VCV := CV.FieldByName("NoWait")
		if VCV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String()+".Config", "NoWait")
		}
		if VCV.Type() != reflect.TypeOf(b) {
			return ErrorMismatchType(V.Elem().Kind().String()+".Config.NoWait", reflect.TypeOf(b), VCV.Type())
		}
		VCV.SetBool(b)
		return nil
	}
}

type ConsumerOption func(interface{}) error

func ConsumerOptionWithQueue(v Queue) ConsumerOption {
	return func(vc interface{}) error {
		V := reflect.ValueOf(vc)
		if V.Kind() != reflect.Pointer {
			return ErrorNotPointer(reflect.ValueOf(vc).Kind())
		}

		CV := V.Elem().FieldByName("Config")
		if CV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String(), "Config")
		}
		VCV := CV.FieldByName("Queue")
		if VCV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String()+".Config", "Queue")
		}
		if VCV.Type() != reflect.TypeOf(v) {
			return ErrorMismatchType(V.Elem().Kind().String()+".Config.Queue", reflect.TypeOf(v), VCV.Type())
		}
		VCV.Set(reflect.ValueOf(v))
		return nil
	}
}

func ConsumerOptionWithName(v string) ConsumerOption {
	return func(vc interface{}) error {
		V := reflect.ValueOf(vc)
		if V.Kind() != reflect.Pointer {
			return ErrorNotPointer(reflect.ValueOf(vc).Kind())
		}

		CV := V.Elem().FieldByName("Config")
		if CV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String(), "Config")
		}
		VCV := CV.FieldByName("Name")
		if VCV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String()+".Config", "Name")
		}
		if VCV.Type() != reflect.TypeOf(v) {
			return ErrorMismatchType(V.Elem().Kind().String()+".Config.Name", reflect.TypeOf(v), VCV.Type())
		}
		VCV.Set(reflect.ValueOf(v))
		return nil
	}
}

func ConsumerOptionWithExclusive(v bool) ConsumerOption {
	return func(vc interface{}) error {
		V := reflect.ValueOf(vc)
		if V.Kind() != reflect.Pointer {
			return ErrorNotPointer(reflect.ValueOf(vc).Kind())
		}

		CV := V.Elem().FieldByName("Config")
		if CV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String(), "Config")
		}
		VCV := CV.FieldByName("Exclusive")
		if VCV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String()+".Config", "Exclusive")
		}
		if VCV.Type() != reflect.TypeOf(v) {
			return ErrorMismatchType(V.Elem().Kind().String()+".Config.Exclusive", reflect.TypeOf(v), VCV.Type())
		}
		VCV.Set(reflect.ValueOf(v))
		return nil
	}
}
func ConsumerOptionWithArgs(v map[string]interface{}) ConsumerOption {
	return func(vc interface{}) error {
		V := reflect.ValueOf(vc)
		if V.Kind() != reflect.Pointer {
			return ErrorNotPointer(reflect.ValueOf(vc).Kind())
		}

		CV := V.Elem().FieldByName("Config")
		if CV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String(), "Config")
		}
		VCV := CV.FieldByName("Args")
		if VCV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String()+".Config", "Args")
		}
		if VCV.Type() != reflect.TypeOf(v) {
			return ErrorMismatchType(V.Elem().Kind().String()+".Config.Args", reflect.TypeOf(v), VCV.Type())
		}
		VCV.Set(reflect.ValueOf(v))
		return nil
	}
}

type ExchangeOption func(interface{}) error

func ExchangeOptionWithName(v string) ExchangeOption {
	return func(vc interface{}) error {
		V := reflect.ValueOf(vc)
		if V.Kind() != reflect.Pointer {
			return ErrorNotPointer(reflect.ValueOf(vc).Kind())
		}

		CV := V.Elem().FieldByName("Config")
		if CV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String(), "Config")
		}
		VCV := CV.FieldByName("Name")
		if VCV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String()+".Config", "Name")
		}
		if VCV.Type() != reflect.TypeOf(v) {
			return ErrorMismatchType(V.Elem().Kind().String()+".Config.Name", reflect.TypeOf(v), VCV.Type())
		}
		VCV.Set(reflect.ValueOf(v))
		return nil
	}
}
func ExchangeOptionWithType(v ExchangeType) ExchangeOption {
	return func(vc interface{}) error {
		V := reflect.ValueOf(vc)
		if V.Kind() != reflect.Pointer {
			return ErrorNotPointer(reflect.ValueOf(vc).Kind())
		}

		CV := V.Elem().FieldByName("Config")
		if CV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String(), "Config")
		}
		VCV := CV.FieldByName("ExchangeType")
		if VCV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String()+".Config", "ExchangeType")
		}
		if VCV.Type() != reflect.TypeOf(v) {
			return ErrorMismatchType(V.Elem().Kind().String()+".Config.ExchangeType", reflect.TypeOf(v), VCV.Type())
		}
		VCV.Set(reflect.ValueOf(v))
		return nil
	}
}

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
func ExchangeOptionAutoDelete(v bool) ExchangeOption {
	return func(vc interface{}) error {
		V := reflect.ValueOf(vc)
		if V.Kind() != reflect.Pointer {
			return ErrorNotPointer(reflect.ValueOf(vc).Kind())
		}

		CV := V.Elem().FieldByName("Config")
		if CV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String(), "Config")
		}
		VCV := CV.FieldByName("AutoDelete")
		if VCV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String()+".Config", "AutoDelete")
		}
		if VCV.Type() != reflect.TypeOf(v) {
			return ErrorMismatchType(V.Elem().Kind().String()+".Config.AutoDelete", reflect.TypeOf(v), VCV.Type())
		}
		VCV.Set(reflect.ValueOf(v))
		return nil
	}
}

func ExchangeOptionDurable(v bool) ExchangeOption {
	return func(vc interface{}) error {
		V := reflect.ValueOf(vc)
		if V.Kind() != reflect.Pointer {
			return ErrorNotPointer(reflect.ValueOf(vc).Kind())
		}

		CV := V.Elem().FieldByName("Config")
		if CV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String(), "Config")
		}
		VCV := CV.FieldByName("Durable")
		if VCV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String()+".Config", "Durable")
		}
		if VCV.Type() != reflect.TypeOf(v) {
			return ErrorMismatchType(V.Elem().Kind().String()+".Config.Durable", reflect.TypeOf(v), VCV.Type())
		}
		VCV.Set(reflect.ValueOf(v))
		return nil
	}
}

// Exchanges declared as `internal` do not accept accept publishings. Internal
// exchanges are useful when you wish to implement inter-exchange topologies
// that should not be exposed to users of the broker.
func ExchangeOptionInternal(v bool) ExchangeOption {
	return func(vc interface{}) error {
		V := reflect.ValueOf(vc)
		if V.Kind() != reflect.Pointer {
			return ErrorNotPointer(reflect.ValueOf(vc).Kind())
		}

		CV := V.Elem().FieldByName("Config")
		if CV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String(), "Config")
		}
		VCV := CV.FieldByName("Internal")
		if VCV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String()+".Config", "Internal")
		}
		if VCV.Type() != reflect.TypeOf(v) {
			return ErrorMismatchType(V.Elem().Kind().String()+".Config.Internal", reflect.TypeOf(v), VCV.Type())
		}
		VCV.Set(reflect.ValueOf(v))
		return nil
	}
}

// A passive exchange is assumed by RabbitMQ to already exist, and attempting to connect to a
// non-existent exchange will cause RabbitMQ to throw an exception. This function
// can be used to detect the existence of an exchange.
func ExchangeOptionPassive(v bool) ExchangeOption {
	return func(vc interface{}) error {
		V := reflect.ValueOf(vc)
		if V.Kind() != reflect.Pointer {
			return ErrorNotPointer(reflect.ValueOf(vc).Kind())
		}

		CV := V.Elem().FieldByName("Config")
		if CV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String(), "Config")
		}
		VCV := CV.FieldByName("Passive")
		if VCV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String()+".Config", "Passive")
		}
		if VCV.Type() != reflect.TypeOf(v) {
			return ErrorMismatchType(V.Elem().Kind().String()+".Config.Passive", reflect.TypeOf(v), VCV.Type())
		}
		VCV.Set(reflect.ValueOf(v))
		return nil
	}
}

type QueueOption func(interface{}) error

type ProducerOption func(interface{}) error

func ProducerOptionOnError(v func(Message)) ProducerOption {
	return func(vc interface{}) error {
		V := reflect.ValueOf(vc)
		if V.Kind() != reflect.Pointer {
			return ErrorNotPointer(reflect.ValueOf(vc).Kind())
		}

		CV := V.Elem().FieldByName("Config")
		if CV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String(), "Config")
		}
		VCV := CV.FieldByName("OnError")
		if VCV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String()+".Config", "OnError")
		}
		if VCV.Type() != reflect.TypeOf(v) {
			return ErrorMismatchType(V.Elem().Kind().String()+".Config.OnError", reflect.TypeOf(v), VCV.Type())
		}
		VCV.Set(reflect.ValueOf(v))
		return nil
	}
}

func ProducerOptionWithRetry(v int) ProducerOption {
	return func(vc interface{}) error {
		V := reflect.ValueOf(vc)
		if V.Kind() != reflect.Pointer {
			return ErrorNotPointer(reflect.ValueOf(vc).Kind())
		}

		CV := V.Elem().FieldByName("Config")
		if CV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String(), "Config")
		}
		VCV := CV.FieldByName("Retry")
		if VCV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String()+".Config", "Retry")
		}
		if VCV.Type() != reflect.TypeOf(v) {
			return ErrorMismatchType(V.Elem().Kind().String()+".Config.Retry", reflect.TypeOf(v), VCV.Type())
		}
		VCV.Set(reflect.ValueOf(v))
		return nil
	}
}

func ProducerOptionWithExchange(v Exchange) ProducerOption {
	return func(vc interface{}) error {
		V := reflect.ValueOf(vc)
		if V.Kind() != reflect.Pointer {
			return ErrorNotPointer(reflect.ValueOf(vc).Kind())
		}

		CV := V.Elem().FieldByName("Config")
		if CV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String(), "Config")
		}
		VCV := CV.FieldByName("Exchange")
		if VCV.Kind() == reflect.Invalid {
			return ErrorMissingFiled(V.Elem().Kind().String()+".Config", "Exchange")
		}
		if VCV.Type() != reflect.TypeOf(v) {
			return ErrorMismatchType(V.Elem().Kind().String()+".Config.Exchange", reflect.TypeOf(v), VCV.Type())
		}
		VCV.Set(reflect.ValueOf(v))
		return nil
	}
}
