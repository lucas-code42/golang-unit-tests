package algoritimos

import (
	"testing"
)

// cenariosTesteSoma:  struct criada para testes da funcao Soma
type cenariosTesteSoma struct {
	primeiroValor int
	segundoValor  int
	valorEsperado int
}

// TestSoma: Implementa a struct cenariosTesteSoma e gera os testes
func TestSoma(t *testing.T) {
	testes := []cenariosTesteSoma{
		{
			primeiroValor: 20,
			segundoValor:  20,
			valorEsperado: 40,
		},
		{
			primeiroValor: 30,
			segundoValor:  50,
			valorEsperado: 80,
		},
	}

	for _, testeAtual := range testes {
		chamandoFuncao := Soma(testeAtual.primeiroValor, testeAtual.segundoValor)
		if chamandoFuncao != testeAtual.valorEsperado {
			t.Error("Erro")
		}
	}
}

// cenariosTesteDivisao: struct criada para testes da funcao Divisao
type cenariosTesteDivisao struct {
	primeiroValor string
	segundoValor  string
	valorEsperado float64
}

// TestDivisao: Implementa a struct cenariosTesteDivisao e gera os testes
func TestDivisao(t *testing.T) {
	testes := []cenariosTesteDivisao{
		{
			primeiroValor: "10",
			segundoValor:  "20",
			valorEsperado: 0.5,
		},
		{
			primeiroValor: "30",
			segundoValor:  "3",
			valorEsperado: 10.0,
		},
		{
			primeiroValor: "50",
			segundoValor:  "10",
			valorEsperado: 5,
		},
		{
			primeiroValor: "0",
			segundoValor:  "1a10",
			valorEsperado: 0,
		},
	}

	for _, teseteAtual := range testes {
		resposta, _ := Divisao(teseteAtual.primeiroValor, teseteAtual.segundoValor)
		if resposta != teseteAtual.valorEsperado {
			t.Error("Erro")
		}
	}
}

// cenariosTesteMultiplicacao: struct criada para testar a funcao Multiplicacao
type cenariosTesteMultiplicacao struct {
	primeiroValor string
	segundoValor  string
	valorEsperado float64
}

// TestMultiplicacao: Implementa a struct cenariosTesteMultiplicacao e gera testes
func TestMultiplicacao(t *testing.T) {
	testes := []cenariosTesteMultiplicacao{
		{
			primeiroValor: "19",
			segundoValor:  "300",
			valorEsperado: 5700,
		},
		{
			primeiroValor: "a999",
			segundoValor:  "100",
			valorEsperado: 0,
		},
		{
			primeiroValor: "299",
			segundoValor:  "1212",
			valorEsperado: 362388,
		},
		{
			primeiroValor: "2.5",
			segundoValor:  "1.3",
			valorEsperado: 3.25,
		},
	}

	for _, testeAtual := range testes {
		chamandoFuncao, _ := Multiplicacao(testeAtual.primeiroValor, testeAtual.segundoValor)
		if chamandoFuncao != testeAtual.valorEsperado {
			t.Error("Erro")
		}
	}
}
