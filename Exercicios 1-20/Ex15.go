package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&x)

	resultado := substituir(x)
	fmt.Println("A string com as vogais substituídas é:", resultado)
}
func substituir(s string) string {
	vogais := "aeiouAEIOU"

	caracteres := strings.Split(s, "")

	for i, char := range caracteres {
		if strings.Contains(vogais, char) {
			caracteres[i] = "*"
		}
	}
	resultado := strings.Join(caracteres, "")

	return resultado
}
