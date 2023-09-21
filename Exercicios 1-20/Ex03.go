package main

import (
	"fmt"
	"strings"
)

func main() {

	var txt, sub, nova string

	fmt.Println("Digite nomes: ")
	fmt.Scan(&txt)

	fmt.Println("Digite o nome a ser removido: ")
	fmt.Scan(&sub)

	fmt.Println("digite o nome a ser adicionado:")
	fmt.Scan(&nova)

	troca := strings.ReplaceAll(txt, sub, nova)
	fmt.Println("Novos nomes: ", troca)
}
