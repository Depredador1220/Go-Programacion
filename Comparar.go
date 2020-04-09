package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hay una señal en la entrada que dice no menores de edad")
	
	var miEdad = 31
	var menorEdad = miEdad < 18
	
	fmt.Printf("Yo tengo %v años, ¿Soy menor de edad? %v\n", miEdad, menorEdad)
}
