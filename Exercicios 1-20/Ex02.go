package main

import (
	"fmt"
	"strings"
)

func main() {

	var str string
	fmt.Println("Digite um texto: ")
	fmt.Scan(&str)

	text := strings.ReplaceAll(str, " ", "")
	fmt.Println("Novo texto sem espaços em branco: ", text)
}
