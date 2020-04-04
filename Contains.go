//Contain

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("¿Te encuentras en una caverna poco iluminada?")
	var comando = "Caminar hacia la salida"
	var salir = strings.Contains(comando, "salida")
	fmt.Println("Salsite de la cueva: ", salir)
}
