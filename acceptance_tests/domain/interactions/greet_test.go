package go_specs_greet_test

import (
	"testing"

	go_specs_greet "github.com/quii/go-specs-greet/domain/interactions"
	"github.com/quii/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(t, specifications.GreetAdapter(go_specs_greet.Greet))
}
