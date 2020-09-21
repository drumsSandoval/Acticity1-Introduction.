package main

import (
	"fmt"
)

func main() {
	var side float64
	fmt.Print("\t\t- Calcular el área del cuadrado\n\n")
	fmt.Print("\tLado: ")
	fmt.Scan(&side)
	result := fmt.Sprintf(" \tEl area es igual a: %.2f", side*side)
	fmt.Println(result)
}
