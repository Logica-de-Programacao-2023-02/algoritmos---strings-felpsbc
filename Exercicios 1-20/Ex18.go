package main

import (
	"fmt"
	"unicode"
)

func main() {
	var x string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&x)

	if Numeros(x) {
		fmt.Println("A string contém apenas números (0 a 9).")
	} else {
		fmt.Println("A string não contém apenas números (0 a 9).")
	}
}
func Numeros(s string) bool {
	for _, char := range s {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}
