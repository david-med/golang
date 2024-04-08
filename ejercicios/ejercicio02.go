package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var err error
var texto string

func Multiplicar() string {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese un númpero a multiplicar")
		if scanner.Scan() {
			numero1, err = strconv.Atoi(scanner.Text())
			if err == nil {
				break
			}
		}
	}
	fmt.Println("Multiplicando ", numero1, "del 1 al 10")
	for i := 1; i <= 10; i++ {
		/*fmt.Println(numero1, "x", i, "=", numero1*i)*/
		texto += fmt.Sprintf("%d x %d = %d \n", numero1, i, numero1*i)
	}
	return texto
}
