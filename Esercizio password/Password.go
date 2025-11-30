package main

import (
	"fmt"
)

func main() {
	var pass string

	fmt.Print("Inserisci password:\n")
	fmt.Scan(&pass)

	if pass != "gosemplice" {
		fmt.Print("Password sbagliata \n")
	} else {
		fmt.Print("Password corretta \n")
	}
}