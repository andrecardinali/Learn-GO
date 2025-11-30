package main

import (
	"fmt"
)

func main() {
	var num1 int
	var num2 int

	fmt.Printf("Inserisci il primo numero: ")
	fmt.Scan(&num1)

	fmt.Printf("Inserisci il secondo numero: ")
	fmt.Scan(&num2)

	if num1 > num2 {
		fmt.Printf("Il numero %d è maggiore di %d\n", num1, num2)
	} else if num1 < num2 {
		fmt.Printf("il numero %d è minore di %d\n", num1, num2)
	} else {
		fmt.Print("I due numeri sono uguali\n")
	}
}
