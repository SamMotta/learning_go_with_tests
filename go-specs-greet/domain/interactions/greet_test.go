package interactions_test

import (
	"testing"

	"github.com/SamMotta/go-specs-greet/domain/interactions"
	"github.com/SamMotta/go-specs-greet/specifications"
	"github.com/alecthomas/assert/v2"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(interactions.Greet),
	)

	t.Run("if name is an empty string default name is 'World'", func(t *testing.T) {
		assert.Equal(t, "Hello, World", interactions.Greet(""))
	})
}
