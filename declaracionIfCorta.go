package main

import (
	"fmt"
	"math/rand"
)

func main() {	
	if num := rand.Intn(3); num == 0 {
		fmt.Println("Aventuras Espaciales")
	} else if num == 1 {
		fmt.Println("Guerras de las Galaxias")
	} else {
		fmt.Println("Invasores del Espacio")
	}
}
