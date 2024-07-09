package ejercicios

import (
	"strconv"
)

func Ejercicio1(text string) (string, int) {
	var texto = ""
	num, err := strconv.Atoi(text)

	if err != nil {
		texto = "Error!!!"
	} else if num <= 100 {
		texto = "El número ingresado es menor o igual a 100 y el número convertido es:"
	} else {
		texto = "El número ingresado es mayor de 100 y el número convertido es:"
	}

	return texto, num
}
