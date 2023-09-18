package main

import "fmt"

func main() {
	var x string
	var reverse string

	fmt.Println("digite uma string: ")
	fmt.Scan(&x)

	for i := len(x) - 1; i >= 0; i-- {
		reverse = reverse + string(x[i])
	}
	if x == reverse {
		fmt.Println("É palidromo ")
	} else {
		fmt.Println("Não é palidromo")
	}
}
