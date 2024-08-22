package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumeros() {

	// Crear un scanner para leer datos del teclado
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese número 1: ")

	// Leer datos del teclado
	if scanner.Scan() {
		// Convertir el texto a un número
		numero1, err = strconv.Atoi(scanner.Text())
		// Validar si hay un error
		if err != nil {
			// Mostrar mensaje de error
			panic("El dato ingresado es incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingrese número 2: ")
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingrese leyenda: ")
	if scanner.Scan() {
		leyenda = scanner.Text()

	}

	fmt.Println(leyenda, numero1*numero2)

}
