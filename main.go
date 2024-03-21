package main

import (
	"fmt"
	"golang/variables"
)

func main() {
	fmt.Println("Hola mundo")
	variables.MostrarEnteros()
	variables.RestoVariables()
	estado, texto := variables.ConviertoaTexto(1549)
	fmt.Println(estado)
	fmt.Println(texto)

}
