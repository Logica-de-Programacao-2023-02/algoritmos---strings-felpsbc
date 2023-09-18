package main

import (
	"fmt"
	"strings"
)

func main() {

	var x string

	fmt.Println("Escreva uma string: ")
	fmt.Scan(&x)

	vogais := []string{"A", "a", "E", "e", "O", "o", "I", "i", "U", "u"}

	for _, vogal := range vogais {
		x = strings.ReplaceAll(x, vogal, "")
	}
	fmt.Println(x)
}
