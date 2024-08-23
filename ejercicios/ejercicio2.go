package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num int
var err error

func TablaDeMultiplicar() {
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

	for i := 1; i <= 10; i++ {
		fmt.Println(num, "x", i, "=", num*i)

	}

}
