package arreglos_slices

import (
	"fmt"
)

var tabla [10]int = [10]int{10, 9, 8}
var matriz [20][30]int

func MuestroArreglos() {
	tabla[7] = 33
	tabla[2] = 20
	tabla2 := [10]string{"David", "Pablo", "Juan"}
	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}
	matriz[7][24] = 15
	fmt.Println(matriz)
}
