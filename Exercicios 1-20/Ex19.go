package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string

	fmt.Println("Escreva uma string: ")
	fmt.Scan(&x)

	reversedString := reverseString(x)

	fmt.Println("A string invertida é:", reversedString)
}

func reverseString(s string) string {
	char := strings.Split(s, "")

	for i, j := 0, len(char)-1; i < j; i, j = i+1, j-1 { // Corrigido aqui
		char[i], char[j] = char[j], char[i]
	}
	reversedString := strings.Join(char, "")
	return reversedString
}
