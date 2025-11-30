package main

import (
	"fmt"
)

func main() {
	var num1 int
	var num2 int
	var result int
	var operator string

	fmt.Print("Inserisci il primo numero:\n")
	fmt.Scan(&num1)

	fmt.Print("Inserisci il secondo numero:\n")
	fmt.Scan(&num2)

	fmt.Print("Che operazione vuoi fare? (+, -, *, /)")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Errore: divisione per zero")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Operatore non valido")
		return
	}

	fmt.Printf("Risultato: %d\n", result)
}
