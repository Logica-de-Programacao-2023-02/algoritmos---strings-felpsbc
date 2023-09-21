package main

import (
	"fmt"
	"strings"
)

func main() {

	var nome string

	fmt.Println("Escreva seu nome em letras minúsculas: ")
	fmt.Scan(&nome)

	Nome := strings.ToUpper(nome)
	fmt.Println("O seu nome em letras maiúsculas ficou assim: ", Nome)
}
