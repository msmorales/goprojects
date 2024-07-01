package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var FechaDeNacimiento time.Time

func RestoVariables() {

	Nombre = "Pedro"
	Estado = true
	Sueldo = 1577.66
	FechaDeNacimiento = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(FechaDeNacimiento)

}

func ConvertoaText(numero int) (bool, string) {

	texto := strconv.Itoa(numero)
	return true, texto
}
