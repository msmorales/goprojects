package main

import (
	"fmt"
	"strings"
)

//"fmt"
//"fmt"
//"github.com/mmorales/goprojects/teclado"
//"github.com/mmorales/goprojects/ejercicios"
//"runtime"
//"github.com/mmorales/goprojects/variables"

func main() {

	// Llamado a la función
	//fmt.Println(variables.ConvertoaText(553))
	//fmt.Println("Hola mundo")
	// OBtiene el sistema operativo
	/*if os := runtime.GOOS; os == "Linux." {
		fmt.Println("Esto no es windows")
	} else {
		fmt.Println("Esto si es windows, es ", os)
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n ", os)

	}*/

	//fmt.Println(ejercicios.Ejercicio1("a"))
	//teclado.IngresoNumeros()
	//iteraciones.Iterar()

	//ejercicios.TablaDeMultiplicar()

	//files.GrabarTabla()

	/*input := []int{3, 4, -7, 5, -6, 2, 5, -1, 8}
	output := SumaACero(input)
	fmt.Println(output)*/ // Debería imprimir: [5, 8]

	// Ejemplo de prueba
	key := "dcj"
	mensaje := "I love prOgrAmming!"
	encriptar := encriptar(key, mensaje)
	fmt.Println(encriptar) // Salida esperada: "dcjI ldcjovdcje prdcjOgrdcjAmmdcjing!"

}

// SumaACero recibe un arreglo de enteros y devuelve un nuevo arreglo que contiene los elementos del arreglo original, pero sin los elementos cuya suma acumulada sea cero.
func SumaACero(arr []int) []int {
	if arr == nil {
		return []int{}
	}

	// Crear un mapa para almacenar la suma acumulativa y la posición y se usa make para inicializar el mapa
	inicial := make(map[int]int)
	// Inicializar la suma acumulativa
	acumulado := 0

	// Iterar sobre el arreglo
	for i := 0; i < len(arr); i++ {
		acumulado += arr[i]

		// Si la suma acumulativa es cero, eliminar los elementos antes del índice actual
		if acumulado == 0 {
			arr = arr[i+1:]
			i = -1 // Reiniciar el índice después de eliminar elementos
			inicial = make(map[int]int)
		} else if pos, exists := inicial[acumulado]; exists {
			// Si la suma acumulativa ya existe en el mapa, eliminar los elementos entre la última aparición y el actual
			arr = append(arr[:pos+1], arr[i+1:]...)
			i = pos
		} else {
			// Almacenar la suma acumulativa y la posición
			inicial[acumulado] = i
		}
	}

	return arr
}

// Función para encriptar el mensaje
func encriptar(key, mensaje string) string {
	// Si el mensaje es nulo o vacío, devolver cadena vacía
	if mensaje == "" {
		return ""
	}

	// Si la clave es nula o vacía, usa "DCJ" como valor predeterminado
	if key == "" {
		key = "DCJ"
	}

	// Definir las vocales
	vowels := "aeiouAEIOU"
	var result strings.Builder

	// Recorrer el mensaje
	for _, ch := range mensaje {
		if strings.ContainsRune(vowels, ch) {
			// Si es vocal, añadir la clave antes del carácter
			result.WriteString(key)
		}
		// Añadir siempre el carácter actual
		result.WriteRune(ch)
	}

	// Devolver el mensaje encriptado
	return result.String()
}
