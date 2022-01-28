package main

import (
	"fmt"
)

func imprimir(message string) {
	fmt.Println("variable... ", message)
}

func variables(a, b int, c string) {
	fmt.Println(a, b, c)
}

func retorno(a int) int {
	return a * 2
}

func dosRetornos(a int) (c, d int) {
	return a + 2, a * 3
}

func main() {
	fmt.Println("variables")
	imprimir("2")
	variables(2, 3, "fa")
	retorno(23)
	v1, v2 := dosRetornos(43)
	v3, _ := dosRetornos(43)
	fmt.Println("v1 y v2 ", v1, v2)
	fmt.Println(v3)
}
