package main

import (
	"fmt"
)

func main() {
	var fh float64
	fmt.Print("\t\t - Convertir de grados Fahrenheit a Celcius (C = (F − 32) * 5/9)\n\n")
	fmt.Print("\tGrados Fahrenheit: ")
	fmt.Scan(&fh)
	celcius := (fh - 32) * 5 / 9
	result := fmt.Sprintf(" \tGrados Celcius: %.2f", celcius)
	fmt.Println(result)
}
