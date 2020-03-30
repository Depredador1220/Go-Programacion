//Viajar a la velocidad de la luz

package main

import (
	"fmt"
)

func main() {
	const velocidadLuz = 299792
	var distancia = 56000000
	
	fmt.Println(distancia/velocidadLuz, " segundos")
	distancia = 401000000
	
	fmt.Println(distancia/velocidadLuz, " segundos")
}
