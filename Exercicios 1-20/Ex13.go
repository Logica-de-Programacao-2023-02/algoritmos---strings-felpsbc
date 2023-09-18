package main

import (
	"fmt"
	"strconv"
)

func main() {

	var x string
	fmt.Println("Digite uma string: ")
	fmt.Scan(&x)

	ints := make([]int, len(x))
	for i, c := range x {
		var err error
		ints[i], err = strconv.Atoi(string(c))
		if err != nil {
			fmt.Println("A sequência é inválida")
			return
		}
	}
	var isCrescente = true
	for i := 0; i < len(ints)-1; i++ {
		atual := ints[i]
		prox := ints[i+1]

		if atual >= prox {
			isCrescente = false
			break
		}
	}
	if isCrescente {
		fmt.Println("É crescente")
	} else {
		fmt.Println("Não é crescente")
	}
}
