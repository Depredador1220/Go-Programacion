package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var contador = 0;
	
	for contador < 10 {
		var num = rand.Intn(10) + 1
		fmt.Println(num)
		
		contador++
	} 
}
