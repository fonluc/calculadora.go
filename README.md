### Calculadora com Testes em Go

Neste exemplo, declaramos duas variáveis **`num1`** e **`num2`** do tipo **`float64`** para armazenar os números fornecidos pelo usuário. Usamos a função **`Scan`** para ler a entrada do usuário e atribuir os valores às variáveis.

Em seguida, realizamos as operações básicas usando os operadores aritméticos **`+`**, **`-`**, **`*`** e **`/`**, e armazenamos os resultados nas variáveis **`soma`**, **`subtracao`**, **`multiplicacao`** e **`divisao`**.

Por fim, imprimimos os resultados na tela.

Ao executar o programa, ele solicitará ao usuário que digite os dois números. Após o usuário inserir os valores e pressionar Enter, o programa calculará a soma, subtração, multiplicação e divisão dos números e exibirá os resultados.

```go
package main

import (
	"fmt"
	"testing"
)

func TestOperacoes(t *testing.T) {
	testes := []struct {
		a, b     float64
		operacao string
		esperado float64
	}{
		{2.5, 3.5, "soma", 6.0},
		{5.0, 2.0, "subtracao", 3.0},
		{2.0, 3.0, "multiplicacao", 6.0},
		{10.0, 2.0, "divisao", 5.0},
	}

	for _, teste := range testes {
		t.Run(teste.operacao, func(t *testing.T) {
			resultado := executarOperacao(teste.a, teste.b, teste.operacao)
			if resultado != teste.esperado {
				t.Errorf("%s incorreta. Resultado: %f, Esperado: %f", teste.operacao, resultado, teste.esperado)
			}
		})
	}
}

func executarOperacao(a, b float64, operacao string) float64 {
	switch operacao {
	case "soma":
		return a + b
	case "subtracao":
		return a - b
	case "multiplicacao":
		return a * b
	case "divisao":
		return a / b
	default:
		panic("Operacao invalida")
	}
}

func main() {
	fmt.Println("Executando testes...")
	// Executa os testes
	resultadoTestes := testing.RunTests(nil, []*testing.M{{Test: TestOperacoes}})
	if !resultadoTestes.Passed() {
		fmt.Println("Alguns testes falharam.")
	} else {
		fmt.Println("Todos os testes foram aprovados.")
	}
}
```

Nessa versão, utilizamos uma estrutura de dados para armazenar os dados dos testes, incluindo os valores de entrada, a operação e o resultado esperado. Em seguida, iteramos sobre essa estrutura de dados usando um loop **`for`** e **`range`** para executar os testes.

Utilizamos a função **`t.Run`** para criar subtestes individualmente nomeados com base na operação. Isso torna mais fácil identificar qual teste falhou, pois o nome do subteste será exibido na saída do teste.

Refatoramos a função **`executarOperacao`** para receber a operação como uma string e usar um **`switch`** para determinar qual operação deve ser executada.
