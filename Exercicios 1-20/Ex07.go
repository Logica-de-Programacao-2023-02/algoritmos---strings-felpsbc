package main

import (
	"fmt"
	"unicode"
)

func main() {

	var x string

	fmt.Println("Digite uma String")
	fmt.Scan(&x)

	numero := false

	for _, char := range x {
		if unicode.IsDigit(char) {
			numero = true
			break
		}
	}
	if numero {
		fmt.Println("A string contem numero")
	} else {
		fmt.Println("A string não contem numero")
	}
}
