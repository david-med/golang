package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var err error

func Multiplicar() {

	scanner := bufio.NewReader(os.Stdin)

	for{
		fmt.Println("Ingrese un númpero a multiplicar")
		if scanner.Scan(){}

			numero1, err = strconv.Atoi(leer.Text())
			if err != nil {
				continue
			}
			for i=1; i<=10; i++{
				fmt.Println(numero1, " x ", i, " = ", numero1*i)
			}
			break
		}
	}
}