package main

import (
	"fmt"
	"os"
)



func main() {
	args := os.Args[1:]
	fmt.Println(args)

	pathENV := os.Getenv("PATH")
	fmt.Println(pathENV)
}
