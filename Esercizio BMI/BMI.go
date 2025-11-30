package main

import (
	"fmt"
)

func main() {
	var peso float64
	var altezza float64
	var bmi float64

	fmt.Print("Inserisci peso in kg: ")
	fmt.Scan(&peso)
	fmt.Print("Inserisci altezza in metri: ")
	fmt.Scan(&altezza)

	bmi = peso / (altezza * altezza)

	switch {
	case (bmi < 18.5):
		fmt.Printf("Il tuo BMI è %.2f, sei sottopeso\n", bmi)
	case (bmi >= 18.5 && bmi < 24.9):
		fmt.Printf("Il tuo BMI è %.2f, sei nella norma\n", bmi)
	case (bmi >= 25 && bmi < 29.9):
		fmt.Printf("Il tuo BMI è %.2f, sei sovrappeso\n", bmi)
	case (bmi >= 30):
		fmt.Printf("Il tuo BMI è %.2f, sei obeso\n", bmi)
	}
}
