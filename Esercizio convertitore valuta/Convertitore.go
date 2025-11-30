package main

import (
	"fmt"
)

func main() {
	var soldi float64
	var valuta1 string
	var valuta2 string
	var EURCHF float64 = 0.93
	var EURUSD float64 = 1.16
	var USDCHF float64 = 0.8

	fmt.Print("Inserisci la somma di denaro da convertire: ")
	fmt.Scan(&soldi)

	fmt.Print("Inserisci la valuta che possiedi (USD, EUR, CHF): ")
	fmt.Scan(&valuta1)

	fmt.Print("Inserisci la valuta in cui vuoi convertire (USD, EUR, CHF): ")
	fmt.Scan(&valuta2)

	switch {
	case valuta1 == "EUR" && valuta2 == "USD":
		fmt.Printf("%.2f Euro sono %.2f Dollari\n", soldi, soldi*EURUSD)
	case valuta1 == "EUR" && valuta2 == "CHF":
		fmt.Printf("%.2f Euro sono %.2f Franchi Svizzeri\n", soldi, soldi*EURCHF)
	case valuta1 == "USD" && valuta2 == "EUR":
		fmt.Printf("%.2f Dollari sono %.2f Euro\n", soldi, soldi/EURUSD)
	case valuta1 == "USD" && valuta2 == "CHF":
		fmt.Printf("%.2f Dollari sono %.2f Franchi Svizzeri\n", soldi, soldi*USDCHF)
	case valuta1 == "CHF" && valuta2 == "EUR":
		fmt.Printf("%.2f Franchi Svizzeri sono %.2f Euro\n", soldi, soldi/EURCHF)
	case valuta1 == "CHF" && valuta2 == "USD":
		fmt.Printf("%.2f Franchi Svizzeri sono %.2f Dollari\n", soldi, soldi/USDCHF)
	default:
		fmt.Println("Conversione non supportata o valute uguali")
	}
}
