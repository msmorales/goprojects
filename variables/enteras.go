package variables

// Importación de paquetes
import "fmt"

// Para que una funcion se encuentre publica debe ir en mayusculas
func MuestroEnteros() {

	// Declaraciónb de variable declarativamente
	var intcomun int

	// Manera compuesta y casteo de tipos
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("intcomun = ", intcomun)
	fmt.Println("intde32 = ", intde32)
	fmt.Println("intde64 = ", intde64)

}
