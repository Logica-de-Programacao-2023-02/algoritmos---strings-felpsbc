package main

import (
	"fmt"
)

func main() {
	var x string

	fmt.Println("Digite uma string numérica: ")
	fmt.Scanln(&x)

	if isDecrescente(x) {
		fmt.Println("É decrescente.")
	} else {
		fmt.Println("Não é decrescente.")
	}
}

func isDecrescente(s string) bool {
	if len(s) <= 1 {
		return false
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] <= s[i+1] {
			return false
		}
	}

	return true
}
