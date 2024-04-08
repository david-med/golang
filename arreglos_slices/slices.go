package arreglos_slices

import (
	"fmt"
)

var tablaS []int = []int{2, 5, 8}
var arreglo [10]int = [10]int{6, 34, 25, 34, 34, 55, 66, 7, 8, 9}

func MuestroSlices() {
	fmt.Println(tablaS)
	porcion1 := arreglo[3:]  //slice creado desde un vector desde la posición 3 hasta el final
	porcion2 := arreglo[:5]  //slice creado desde un vector desde la posición 0 hasta la 5
	porcion3 := arreglo[4:7] //slice creado desde un vector desde la posición 4 hasta la 7

	fmt.Println(porcion1)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))

}
