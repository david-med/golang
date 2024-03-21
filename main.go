package main

import (
	"fmt"
	"golang/variables"
	"runtime"
)

func main() {
	fmt.Println("Hola mundo")
	variables.MostrarEnteros()
	variables.RestoVariables()
	estado, texto := variables.ConviertoaTexto(1549)
	fmt.Println(estado)
	fmt.Println(texto)
	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es windows, es: ", os)
	} else {
		fmt.Println("Esto es windows, ", os)
	}

	switch os := runtime.GOOS; os {
	case "Linux.":
		fmt.Println("Esto es: Linux, ", os)
	case "darwin":
		fmt.Println("Esto es: darwin, ", os)
	case "windows":
		fmt.Println("Esto es: windows, ", os)
	default:
		fmt.Printf("%s \n", os)
	}
}
