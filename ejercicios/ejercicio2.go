package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num int
var err error
var texto string

func TablaDeMultiplicar() string {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingrese número: ")
		// Leer datos del teclado
		if scanner.Scan() {
			// Convertir el texto a un número
			num, err = strconv.Atoi(scanner.Text())
			// Validar si hay un error
			if err != nil {
				// Mostrar mensaje de error
				continue
				//panic("El dato ingresado es incorrecto" + err.Error())
			} else {
				break
			}
		}
	}

	fmt.Println("Tabla de multiplicar del", num)
	// Tabla de multiplicar
	for i := 1; i <= 10; i++ {

		// Imprimir la tabla de multiplicar
		//fmt.Println(num, "x", i, "=", num*i)
		// Guardar la tabla de multiplicar en una variable
		texto += fmt.Sprintln(num, "x", i, "=", num*i)

	}

	return texto
}
