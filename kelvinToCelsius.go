package main

import "fmt"

const kelvinBase float64 = 273.15

func main() {
	kelvinBoiling := 373.2
	fmt.Printf("Ponto de ebulição da água: %.2fK (Kelvin), %.2fC (Celsius)", kelvinBoiling, kelvinBoiling-kelvinBase)
}
