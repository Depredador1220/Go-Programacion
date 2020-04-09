package main

import (
	"fmt"
)

func main() {
	fmt.Println("El año actual es 2100, ¿Sera año bisiesto?")
	
	var anio     = 2100
	var bisiesto = anio % 400 == 0 || (anio % 4 == 0 || anio % 100 != 0) 
	
	if bisiesto {
		fmt.Println("Es año bisiesto")
	} else {
		fmt.Println("No es año bisisesto")
	}
}
