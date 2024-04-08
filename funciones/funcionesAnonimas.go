package funciones

func Calculos() {

	var numero3 int = 10
	var numero4 int = 100

	calculo := func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero3
	}
	println(calculo(10, 25))

	calculo = func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero4
	}
	println(calculo(10, 25))

}
