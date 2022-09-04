package algoritimos

import (
	"errors"
	"strconv"
)

// Soma: Retorna soma entre dois numeros
func Soma(n1, n2 int) int {
	return n1 + n2
}

// Divisao: Verifica se os parametros passados sao numeros, se forem retorna uma divisao e um erro nil, caso contrario um 0 e um erro
func Divisao(n1, n2 string) (float64, error) {
	validate := isFloat(n1, n2)
	if !validate {
		return 0, errors.New("verifique os parametros enviados para esta funcao")
	}

	value1, _ := strconv.ParseFloat(n1, 64)
	value2, _ := strconv.ParseFloat(n2, 64)

	divisaoPorZero := impossivelDivisaoPorZero(value1, value2)
	if !divisaoPorZero {
		return 0, errors.New("impossivel divisao por zero")
	}

	calc := value1 / value2
	return calc, nil
}

// Multiplicacao: Verifica se os parametros passados sao numeros, se forem retorna uma Multiplicacao e um erro nil, caso contrario um 0 e um erro
func Multiplicacao(n1, n2 string) (float64, error) {
	validate := isFloat(n1, n2)
	if !validate {
		return 0, errors.New("verifique os parametros enviados para esta funcao")
	}

	value1, _ := strconv.ParseFloat(n1, 64)
	value2, _ := strconv.ParseFloat(n2, 64)
	calc := value1 * value2
	return calc, nil
}

// isfloat: Verifica se os dois números passados podem ser convertidos para float, se sim retorna true contrártio retorna false
func isFloat(n1, n2 string) bool {
	var verificacaoN1 bool
	_, err1 := strconv.ParseFloat(n1, 64)
	if err1 == nil {
		verificacaoN1 = true
	}

	var verificacaoN2 bool
	_, err2 := strconv.ParseFloat(n2, 64)
	if err2 == nil {
		verificacaoN2 = true
	}

	if !verificacaoN1 || !verificacaoN2 {
		return false
	} else {
		return true
	}
}

// Verifica se os valores passados e convertidos para float dentro de uma divisao são 0
func impossivelDivisaoPorZero(n1, n2 float64) bool {
	if n1 == 0 || n2 == 0 {
		return false
	}
	return true
}
