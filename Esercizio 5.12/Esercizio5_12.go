package main

import "fmt"

func main() {
	LOOP:
		for i := 0; i <= 10; i++ {
			if i == 5 {
				//break LOOP
				//continue LOOP
				goto LOOP
			}
		fmt.Println("Value of i: ", i)
		}
}