package main

import (
	"fmt"
	"unicode"
)

func main() {
	var x string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&x)

	Case, numPalavras := verificaCase(x)

	if Case {
		fmt.Printf("A string está em camelCase e possui %d palavra(s).\n", numPalavras)
	} else {
		fmt.Println("A string não está em camelCase.")
	}
}
func verificaCase(s string) (bool, int) {
	Case := true
	numPalavras := 1

	for i, char := range s {
		if !unicode.IsLetter(char) {
			Case = false
			break
		}

		if i > 0 && unicode.IsUpper(char) {
			numPalavras++
		}
	}

	return Case, numPalavras
}
