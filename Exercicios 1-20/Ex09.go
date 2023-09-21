package main

import (
	"fmt"
	"strings"
)

func main() {

	var x, letra, sub string

	fmt.Println("Digite uma string: ")
	fmt.Scan(&x)

	fmt.Println("Digite a letra que deseja trocar: ")
	fmt.Scan(&letra)

	fmt.Println("Digite a letra substituta: ")
	fmt.Scan(&sub)

	resultado := substituirLetra(x, letra, sub)

	fmt.Println("A string após a modificação é: ", resultado)
}

func substituirLetra(s, alvo, susbstituta string) string {
	resultado := strings.ReplaceAll(s, alvo, susbstituta)
	return resultado
}
