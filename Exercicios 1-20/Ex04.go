package main

import "fmt"

func main() {

	var x, y string

	fmt.Println("Digite uma string")
	fmt.Scan(&x)

	fmt.Println("Digite outra string")
	fmt.Scan(&y)

	if x == y {
		fmt.Println("As strings são iguais")
	} else {
		fmt.Println("As strings são diferentes")
	}

}
