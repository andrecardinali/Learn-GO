package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		original := i & 0xFF
		complemento := ^i & 0xFF

		fmt.Printf("Numero: %2d -> Bin: %08b | Complemento: %08b\n", original, original, complemento)
	}
}