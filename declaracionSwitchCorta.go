package main

import (
	"fmt"
	"math/rand"
)

func main() {	
	switch num := rand.Intn(10); num {
		case 0 :
			fmt.Println("Aventuras Espaciales")
		case 1 : 
			fmt.Println("Guerras de las Galaxias")
		case 2 :
			fmt.Println("Invasores del Espacio")
		default:
			fmt.Println("Otro resultado #", num)
	}
}
