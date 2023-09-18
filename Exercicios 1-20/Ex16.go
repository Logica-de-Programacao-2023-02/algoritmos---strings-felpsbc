package main

import (
	"fmt"
	"strings"
)

func main() {

	var x, y string

	fmt.Println("Digite duas Strings: ")
	fmt.Scan(&x, &y)

	if strings.Contains(x, y) {
		fmt.Println("A substring esta contida")
	} else {
		fmt.Println("A substring não está contida")
	}
}
