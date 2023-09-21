package main

import (
	"fmt"
	"strings"
)

func main() {

	var x string

	fmt.Println("Digite uma string: ")
	fmt.Scan(&x)

	p := strings.Fields(x)
	contador := len(p)

	fmt.Printf("A string têm %d palavras.\n", contador)
}
