package main

import (
	"fmt"
)

func main() {
	var (
		base   = 0.0
		height = 0.0
	)

	fmt.Print("\t\t- Calcular el área de un Triangulo\n\n")
	fmt.Print("\tbase: ")
	fmt.Scan(&base)
	fmt.Print("\tAltura: ")
	fmt.Scan(&height)
	result := fmt.Sprintf(" \tEl area es igual a: %.2f", (base*height)/2)
	fmt.Println(result)
}
