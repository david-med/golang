package ejercicios

import "strconv"

func TextoAValor(texto string) (int, string) {
	var mensaje string
	entero, err := strconv.Atoi(texto)
	if err == nil {
		var mensaje string
		if entero > 100 {
			mensaje = "Es mayor a 100"
		} else {
			mensaje = "Es menor a 100"
		}
		return entero, mensaje
	} else {
		mensaje = "Se presentó error en la conversión del string " + err.Error()
	}
	return 0, mensaje
}
