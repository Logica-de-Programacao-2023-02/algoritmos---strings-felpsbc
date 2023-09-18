package main

import "fmt"

func main() {
	var x string

	fmt.Println("digite uma string: ")
	fmt.Scan(&x)

	var unicas []string

	for i := 0; i < len(x); i++ {
		isUnique := true
		for j := 1; j < len(x); j++ {
			if i != j && x[i] == x[j] {
				isUnique = false
				break
			}
		}
		if isUnique {
			unicas = append(unicas, string(x[i]))
		}
	}
	fmt.Println(unicas)
}
