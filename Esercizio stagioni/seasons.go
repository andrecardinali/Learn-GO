package main

import "fmt"

func main() {
	var mese string

	fmt.Print("Inserisci il mese dell'anno: ")
	fmt.Scan(&mese)

	switch mese {
	case "dicembre", "gennaio", "febbraio":
		fmt.Println("La stagione è Inverno")
	case "marzo", "aprile", "maggio":
		fmt.Println("La stagione è Primavera")
	case "giugno", "luglio", "agosto":
		fmt.Println("La stagione è Estate")
	case "settembre", "ottobre", "novembre":
		fmt.Println("La stagione è Autunno")
	default:
		fmt.Println("Mese non valido")
	}
}
