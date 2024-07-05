package specifications

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

type Greeter interface {
	Greet(string) (string, error)
}

func GreetSpecification(t testing.TB, greeter Greeter) {
	got, err := greeter.Greet("Mike")
	assert.NoError(t, err)
	assert.Equal(t, got, "Hello, Mike!")
}

type MeanGreeter interface {
	Curse(name string) (string, error)
}

func CurseSpecification(t testing.TB, meany MeanGreeter) {
	got, err := meany.Curse("Chris")
	assert.NoError(t, err)
	assert.Equal(t, got, "Go to hell, Chris!")
}
