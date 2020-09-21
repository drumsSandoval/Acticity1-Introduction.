package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64
	fmt.Print("\t\t- Calcular el área de un Circulo\n\n")
	fmt.Print("\tradio: ")
	fmt.Scan(&radius)

	result := fmt.Sprintf(" \tEl area es igual a: %.2f", math.Pi*(radius*radius))
	fmt.Println(result)
}
