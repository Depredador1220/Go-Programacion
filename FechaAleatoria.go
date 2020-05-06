package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {
	año      := 2020
	mes      := rand.Intn(12) + 1
	diaEnMes := 31	

	switch mes {
		case 2:
			diaEnMes = 28
		case 4, 6, 9, 11:
			diaEnMes = 30
	}
	
	dia := rand.Intn(diaEnMes) + 1
	fmt.Println(era, año, mes, dia)
}
