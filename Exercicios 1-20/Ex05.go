package main

import (
	"fmt"
	"strconv"
)

func main() {

	var x string
	fmt.Println("Digite um número:")
	fmt.Scan(&x)

	_, err := strconv.ParseFloat(x, 64)
	if err == nil {
		fmt.Println("É um número flutuante")
	} else {
		fmt.Println("O valor não é valido em float")
	}
}
