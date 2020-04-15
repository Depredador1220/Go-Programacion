package main

import (
	"fmt"
	"time"
)

func main() {
	var contador = 10
	
	for contador > 0 {
		fmt.Println(contador)
		time.Sleep(time.Second)
		contador--
	}		
	
	fmt.Println("Termina")	
}
