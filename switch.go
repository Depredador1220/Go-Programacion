package main

import (
	"fmt"
)

func main() {
	fmt.Println("Que camino tomarias")
	
	var comando = "Leer señal"
	
	switch comando {
		case "Ve al este":
			fmt.Println("Llegaras a la montaña")
		
		case "Entrar a la cueva", "Ve adentro":
			fmt.Println("Encontraste una caverna")
			
		case "Leer señal":
			fmt.Println("La señal dice 'peligro'")
			
		default: 
			fmt.Println("No sucedio nada")		
		
	}
}
