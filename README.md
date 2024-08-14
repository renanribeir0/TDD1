<img width="575" alt="Captura de tela 2024-08-13 102947" src="https://github.com/user-attachments/assets/191c1907-75d0-43f0-9596-f02a1a4c2d51">

## TestSayHello

--- PASS: TestSayHello (0.00s)

## TestOddOrEven

=== RUN   TestOddOrEven
=== RUN   TestOddOrEven/Check_Non_Negative_Numbers
--- PASS: TestOddOrEven/Check_Non_Negative_Numbers (0.00s)
=== RUN   TestOddOrEven/Check_Negative_Numbers
--- PASS: TestOddOrEven/Check_Negative_Numbers (0.00s)
--- PASS: TestOddOrEven (0.00s)

```markdown
# go_test_starter

Este projeto é um exemplo de como escrever e executar testes em Go utilizando o framework de testes `testify`.

## Estrutura do Projeto

- `starter.go`: Contém as funções `SayHello` e `OddOrEven`.
- `starter_test.go`: Contém os testes para as funções `SayHello` e `OddOrEven`.

## Funções

### SayHello

```go
func SayHello(name string) string
```

Retorna uma saudação personalizada para o nome fornecido.

### OddOrEven

```go
func OddOrEven(num int) string
```

Verifica se um número é par ou ímpar, incluindo suporte para números negativos.

## Testes

Os testes são escritos usando o pacote `testify/assert`.

### TestSayHello

Verifica se a função `SayHello` retorna a saudação correta para o nome fornecido.

### TestOddOrEven

Verifica se a função `OddOrEven` retorna corretamente se os números, tanto positivos quanto negativos, são pares ou ímpares.

## Executando os Testes

Para executar os testes, você precisará do Go instalado em sua máquina. Siga os passos abaixo:

1. Clone o repositório:
   ```sh
   git clone https://github.com/renanribeir0/TDD1.git
   ```

2. Navegue até o diretório do projeto:
   ```sh
   cd TDD1/go_test_starter
   ```

3. Instale as dependências:
   ```sh
   go get -t ./...
   ```

4. Execute os testes:
   ```sh
   go test -v
   ```

## Dependências

Este projeto utiliza o pacote `testify` para testes. Para instalá-lo, execute:

```sh
go get github.com/stretchr/testify
```

## Contribuições

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou pull requests.


# Go Test Tutorial

Este repositório contém um tutorial de TDD em Go utilizando a biblioteca testify.

## Passos para Executar

1. Clone o repositório:
   ```sh
   git clone https://github.com/seu-usuario/go_test_tutorial.git
   cd go_test_tutorial
