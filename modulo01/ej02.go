package main

import "fmt"

func main() {
	const pi float64 = 3.24
	const pi2 = 3.13

	fmt.Println(pi, pi2)
	//variables
	base := 12
	var altura int = 14
	var area int
	area = base * altura
	fmt.Println(area)
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)
	const baseCuadrado = 20
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println(areaCuadrado)
}
