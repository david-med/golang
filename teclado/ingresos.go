package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error
var hola string

func IngresoNumeros() {

	leer := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese número 1: ")
	if leer.Scan() {

		numero1, err = strconv.Atoi(leer.Text())
		if err != nil {
			panic("el dato ingresado es incorrecto" + err.Error())
		}
	}
	hola = "otra vez"
	fmt.Println("Ingrese número 2: ")
	if leer.Scan() {
		numero2, err = strconv.Atoi(leer.Text())
		if err != nil {
			panic("el dato ingresado es incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingrese Leyenda: ")
	if leer.Scan() {
		leyenda = leer.Text()
	}

	println(leyenda, numero1*numero2, hola)
}

func Hola() {
	leyenda = "hola"
	println(leyenda)
}
