package files

import (
	"bufio"
	"fmt"
	"golang/ejercicios"
	"io/ioutil"
	"os"
)

var fileName string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.Multiplicar()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("error creando archivo" + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.Multiplicar()
	//if aAppend(fileName, texto) == false {
	if !aAppend(fileName, texto) {
		fmt.Println("Error al grabar en el archivo")
	}
}

func aAppend(file string, texto string) bool {
	arch, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error abriendo archivo para escribir" + err.Error())
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error escribiendo archivo " + err.Error())
		return false
	}
	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error al leer el archivo" + err.Error())
		return
	}
	fmt.Println(string(archivo))
}

func LeoArchivo2() {
	archivo, err := os.Open(fileName)
	if err != nil {
		println("Error leyendo archivo" + err.Error())
		return
	}
	var i int = 1
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> ", i, registro)
		i++
	}
	archivo.Close()
}
