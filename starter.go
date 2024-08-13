package starter

import (
	"fmt"
	"math"
)

// SayHello retorna uma saudação personalizada.
// Utiliza a função fmt.Sprintf para formatar a string.
func SayHello(name string) string {
	return fmt.Sprintf("Hello %v. Welcome!", name)
}

// OddOrEven verifica se um número é par ou ímpar.
// Utiliza a função math.Mod para calcular o resto da divisão por 2.
// Adiciona suporte para números negativos.
func OddOrEven(num int) string {
	criteria := math.Mod(float64(num), 2)
	if criteria == 1 || criteria == -1 {
		return fmt.Sprintf("%v is an odd number", num)
	}
	return fmt.Sprintf("%v is an even number", num)
}
