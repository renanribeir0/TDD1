package starter_test

import (
	"testing"

	starter "github.com/renanribeir0/go_test_tutorial/go_test_starter"
	"github.com/stretchr/testify/assert"
)

func TestSayHello(t *testing.T) {
	greeting := starter.SayHello("William")
	assert.Equal(t, "Hello William. Welcome!", greeting)

	another_greeting := starter.SayHello("asdf ghjkl")
	assert.Equal(t, "Hello asdf ghjkl. Welcome!", another_greeting)
}

func TestOddOrEven(t *testing.T) {
	t.Run("Check Non Negative Numbers", func(t *testing.T) {
		assert.Equal(t, "45 is an odd number", starter.OddOrEven(45))
		assert.Equal(t, "42 is an even number", starter.OddOrEven(42))
		assert.Equal(t, "0 is an even number", starter.OddOrEven(0))
	})
	t.Run("Check Negative Numbers", func(t *testing.T) {
		assert.Equal(t, "-45 is an odd number", starter.OddOrEven(-45))
		assert.Equal(t, "-42 is an even number", starter.OddOrEven(-42))
	})
}
