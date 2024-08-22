package iteraciones

import "fmt"

func Iterar() {

	for i := 5; i <= 100; i += 5 {

		// Salta a la siguiente iteración
		/*if i == 40 {
			break
		}*/
		// No se ejecuta la iteración actual
		if i == 40 {
			continue
		}

		fmt.Println(i)

	}

}
