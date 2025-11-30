package main

import (
	"fmt"
)

func main() {
	var num int
	
	fmt.Printf("Inserisci un numero: \n")
	fmt.Scan(&num)

	if num %2 == 0 {
		fmt.Printf("Il numero %d è pari \n", num)
	} else {
		fmt.Printf("Il numero %d è dispari \n", num)
	}
}