//Programa piensa un numero del 1 al 10

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num = rand.Intn(10) + 1
	fmt.Println(num)
	
	num = rand.Intn(10) + 1
	fmt.Println(num)
}
