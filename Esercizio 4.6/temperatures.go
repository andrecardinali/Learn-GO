package main

import "fmt"

func main() {
	var c float64
	var f float64
	var scelta int

	fmt.Println("Scegli la conversione:")
	fmt.Println("Scegli 1 per convertire Celsius in Fahrenheit")
	fmt.Println("Scegli 2 per convertire Fahrenheit in Celsius")
	fmt.Scanln(&scelta)

	if scelta == 1 {
		fmt.Print("Inserisci la temperatura in Celsius: ")
		fmt.Scanln(&c)
		fmt.Println("La temperatura in Fahreinheit è: ", cToF(c))
	} else if scelta == 2 {
		fmt.Print("Inserisci la temperatura in Fahreinheit: ")
		fmt.Scanln(&f)
		fmt.Println("La temperatura in Celsius è: ", fToC(f))
	} else {
		fmt.Println("Scelta non valida")
	}
}

func cToF(c float64) float64 {
	return c*9.0/5.0 + 32.0
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
