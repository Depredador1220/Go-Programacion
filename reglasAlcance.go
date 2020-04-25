package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {

	año := 2020 	
	
	switch mes := rand.Intn(12) + 1; mes {
		case 2 :
			dia := rand.Intn(28) + 1
			fmt.Println(era, año, mes, dia)
		case 4, 6, 9, 11 : 
			dia := rand.Intn(30) + 1
			fmt.Println(era, año, mes, dia)
		default:
			dia := rand.Intn(31) + 1
			fmt.Println(era, año, mes, dia)
	}
}
