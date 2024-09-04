package files

import (
	//	"bufio"
	"fmt"
	//	"ioutil"
	"os"

	"github.com/mmorales/goprojects/ejercicios"
)

var fileName string = "./files/txt/tabla.txt"

func GrabarTabla() {
	// Llamado a la función de la tabla de multiplicar
	var texto string = ejercicios.TablaDeMultiplicar()
	archivo, err := os.Create(filename)

	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()

}

func SumaTabla() {
	// Llamado a la función de la tabla de multiplicar
	var texto string = ejercicios.TablaDeMultiplicar()
	if !Append(fileName, texto) {
		fmt.Println("Error al concatenar texto al archivo")
	}

}

func Append(filen string, texto string) bool {
	// Se abre el archivo en modo de escritura y se le da permisos de lectura y escritura
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		// Si hay un error se imprime el mensaje y se retorna false
		fmt.Println("Error durante el Append" + err.Error())
		return false
	}

	// Se escribe el texto en el archivo
	_, err = arch.WriteString(texto)
	if err != nil {
		// Si hay un error se imprime el mensaje y se retorna false
		fmt.Println("Error durante el WhiteString" + err.Error())
		return false

	}

	// Se cierra el archivo
	arch.Close()

	/*func LeoArchivo () {
		archivo , err := ioutil.ReadFile(fileName)
	}*/

}
