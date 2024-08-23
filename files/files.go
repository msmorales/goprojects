package files

import (
	//	"bufio"
	"fmt"
	//	"ioutil"
	"os"

	"github.com/mmorales/goprojects/ejercicios"
)

func GrabarTabla() {
	// Llamado a la función de la tabla de multiplicar
	var texto string = ejercicios.TablaDeMultiplicar()
	archivo, err := os.Create("./files/txt/tabla.txt")

	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()

}

/*
func SumaTabla (){
	// Llamado a la función de la tabla de multiplicar
	var texto string = ejercicios.TablaDeMultiplicar()

}*/
