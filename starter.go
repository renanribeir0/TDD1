package starter

import "fmt"

// SayHello retorna uma saudação para a pessoa nomeada.
func SayHello(name string) string {
	return fmt.Sprintf("Hello %v. Welcome!", name)
}

// OddOrEven determina se um número é ímpar ou par.
// Utiliza a operação de módulo para verificar.
func OddOrEven(num int) string {
	criteria := num % 2
	if criteria == 1 || criteria == -1 {
		return fmt.Sprintf("%v is an odd number", num)
	}
	return fmt.Sprintf("%v is an even number", num)
}
