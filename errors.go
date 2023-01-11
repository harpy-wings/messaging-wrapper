package messagingwrapper

import (
	"fmt"
	"reflect"
)

var (
	ErrorNotPointer = func(t reflect.Kind) error { return fmt.Errorf("base object is %s not a pointer", t) }

	ErrorMissingFiled = func(base, filedName string) error {
		return fmt.Errorf("base object %s has no filed name %s", base, filedName)
	}

	ErrorMismatchType = func(base string, expected, got reflect.Type) error {
		return fmt.Errorf("base %s has a wrong type, expected %s got %s", base, expected, got)
	}

	ErrorNotFound = func(args ...interface{}) error { return fmt.Errorf("item not found %v", args) }
	
)
