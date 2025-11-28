package main

import (
	"fmt"
	"os"
)

func main() {
	sysname, _ := os.Hostname()
	fmt.Printf("The name of your system is: %s\n", sysname)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}
