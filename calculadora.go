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
