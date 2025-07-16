package main

import (
	"fmt"
	"os"
	"strings"
)



func main() {
	args := os.Args[1:]

	pathENV := os.Getenv("PATH")
	pathArr := strings.Split(pathENV, ":")
	// fmt.Println(pathArr)

	var argsPath []string

	argLen := len(args)
	for i := 0; i < argLen; i++ {
		for _, path := range pathArr {
			fullPath := path + "/" + args[i]
			if _, err := os.Stat(fullPath); err == nil {
				argsPath = append(argsPath, fullPath)
				break
			}
		}
	}

	for _, path := range argsPath {
		fmt.Println(path)
	}
}
