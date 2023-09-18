package main

import "fmt"

func main() {
	var x, y string
	fmt.Println("Escreva duas palavras: ")
	fmt.Scan(&x, &y)

	xMapa := make(map[rune]int)
	yMapa := make(map[rune]int)

	for _, c := range x {
		xMapa[c]++
	}
	for _, c := range y {
		yMapa[c]++
	}

	for char, qtdX := range xMapa {
		qtdY := yMapa[char]
		if qtdX != qtdY {
			fmt.Println("As palavras não são anagramas")
			return
		}
	}

	fmt.Println(xMapa)
	fmt.Println(yMapa)
	fmt.Println("As palavras são anagramas")
}
