package main

import (
	"fmt"
	"unit-tests/algoritimos"
)

func main() {
	fmt.Print("Soma: ")
	fmt.Println(algoritimos.Soma(10, 20))
	
	fmt.Print("Divisao: ")
	fmt.Println(algoritimos.Divisao("2", "0"))

	fmt.Print("Multiplicação: ")
	fmt.Println(algoritimos.Multiplicacao("101", "202"))
	
}
