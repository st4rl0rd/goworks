package main

import (
	"fmt"
)

func main() {
	varMessage := "hola... "
	secMessage := "wil"
	fmt.Println(varMessage, secMessage)
	nombre := "DEBCur"
	cursos := 400
	fmt.Printf("%s tiene %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene %v cursos\n", nombre, cursos)

	message := fmt.Sprintf("%v tiene %v cursos\n", nombre, cursos)
	fmt.Println(message)

	fmt.Printf("varMessage: %T", varMessage)
	fmt.Printf("cursos: %T", cursos)
}
