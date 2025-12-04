package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Repeat("*", 20))
    for i := 0; i < 6; i++ {
		fmt.Print("*                  *\n")
	}
	fmt.Println(strings.Repeat("*", 20))
}