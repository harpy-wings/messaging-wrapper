package messagingwrapper

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBindOptionWithArgs(t *testing.T) {
	value := map[string]interface{}{"key1": "value1", "key2": true}
	t.Run("Success", func(t *testing.T) {
		var Case struct {
			Config struct {
				Args map[string]interface{}
			}
		}

		err := BindOptionWithArgs(value)(&Case)
		require.NoError(t, err)
		require.Len(t, Case.Config.Args, 2)
	})

	t.Run("Fail/Invalid Type", func(t *testing.T) {
		err := BindOptionWithArgs(value)("Unexpected Type")
		require.Error(t, err)
	})
	t.Run("Fail/Missing Filed", func(t *testing.T) {
		var Case struct {
		}
		err := BindOptionWithArgs(value)(&Case)
		require.Error(t, err)
	})
	t.Run("Fail/Missing Filed", func(t *testing.T) {
		var Case struct {
			Config struct {
			}
		}
		err := BindOptionWithArgs(value)(&Case)
		require.Error(t, err)
	})

	t.Run("Fail/Mismatch Type", func(t *testing.T) {
		var Case struct {
			Config struct {
				Args string
			}
		}
		err := BindOptionWithArgs(value)(&Case)
		require.Error(t, err)
	})

}

func TestBindOptionWithNoWait(t *testing.T) {
	value := true
	t.Run("Success", func(t *testing.T) {
		var Case struct {
			Config struct {
				NoWait bool
			}
		}
		err := BindOptionWithNoWait(value)(&Case)
		require.NoError(t, err)
		require.Equal(t, Case.Config.NoWait, value)
	})

	t.Run("Fail/Invalid Type", func(t *testing.T) {
		err := BindOptionWithNoWait(value)("Unexpected Type")
		require.Error(t, err)
	})
	t.Run("Fail/Missing Filed", func(t *testing.T) {
		var Case struct {
		}
		err := BindOptionWithNoWait(value)(&Case)
		require.Error(t, err)
	})
	t.Run("Fail/Missing Filed", func(t *testing.T) {
		var Case struct {
			Config struct {
			}
		}
		err := BindOptionWithNoWait(value)(&Case)
		require.Error(t, err)
	})

	t.Run("Fail/Mismatch Type", func(t *testing.T) {
		var Case struct {
			Config struct {
				NoWait string
			}
		}
		err := BindOptionWithNoWait(value)(&Case)
		require.Error(t, err)
	})
}

func TestProducerOptionOnError(t *testing.T) {
	value := func(msg Message) { log.Printf("%v cahced ;)\n", msg) }
	t.Run("Success", func(t *testing.T) {
		var Case struct {
			Config struct {
				OnError func(msg Message)
			}
		}
		err := ProducerOptionOnError(value)(&Case)
		require.NoError(t, err)
		Case.Config.OnError(nil)
	})

	t.Run("Fail/Invalid Type", func(t *testing.T) {
		err := ProducerOptionOnError(value)("Unexpected Type")
		require.Error(t, err)
	})
	t.Run("Fail/Missing Filed", func(t *testing.T) {
		var Case struct {
		}
		err := ProducerOptionOnError(value)(&Case)
		require.Error(t, err)
	})
	t.Run("Fail/Missing Filed", func(t *testing.T) {
		var Case struct {
			Config struct {
			}
		}
		err := ProducerOptionOnError(value)(&Case)
		require.Error(t, err)
	})

	t.Run("Fail/Mismatch Type", func(t *testing.T) {
		var Case struct {
			Config struct {
				NoWait string
			}
		}
		err := ProducerOptionOnError(value)(&Case)
		require.Error(t, err)
	})
}
