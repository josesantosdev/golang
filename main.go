package main

import (
	"fmt"
	"projetoTeste/calc"
)

func main() {
	var num int = 4
	num2 := 0

	result, err := calc.Dividir(num, num2)

	if err != nil {
		fmt.Println("Erro", err)
		return
	}

	fmt.Println(result)
}
