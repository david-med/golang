package main

import (
	"golang/files"
)

/*
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
*/
func main() {

	/*val, msg := ejercicios.TextoAValor("200")
	fmt.Println(val, msg)

	teclado.IngresoNumeros()
	iteraciones.Iteraciones()*/
	//fmt.Println(ejercicios.Multiplicar())
	//files.GrabaTabla()
	//files.SumaTabla()
	//files.LeoArchivo()
	files.LeoArchivo2()
}
