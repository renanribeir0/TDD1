package starter_test

import (
	"testing"

	starter "github.com/renanribeir0/TDD1/go_test_starter"
	"github.com/stretchr/testify/assert"
)

// TestSayHello testa a função SayHello.
// Verifica se a saudação retornada está correta.
func TestSayHello(t *testing.T) {
	greeting := starter.SayHello("Renan")
	assert.Equal(t, "Hello Renan. Welcome!", greeting)
}

// TestOddOrEven testa a função OddOrEven.
// Divide os testes em subtestes para números não negativos e negativos.
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
