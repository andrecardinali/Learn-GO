package main

import (
	"fmt"
	"strings"
)

func main() {
	for i := 1; i <25; i++ {
		fmt.Println(strings.Repeat("G", i))
	}
	fmt.Print("\n")
	fmt.Print("Altra soluzione:\n")
	fmt.Print("\n")

	for l := 0; l < 25; l++ {
		for j := 0; j <= l; j++ {
			fmt.Print("G")
		}
		fmt.Print("\n")
	}
}