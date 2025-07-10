package main

import (
	"fmt"
	"log"
	"os"
)



func main() {

	if(len(os.Args) < 3) {
		log.Fatal("not enough arguments")
	}

	flag := os.Args[1]
	fileName := os.Args[2]

	switch flag{
	case "-c": 
		content, err := os.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)	
		}
		fmt.Println(len(content), fileName)
	case "-l":
		content, err := os.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)	
		}
		count := 0
		for _, ch := range content {
			if ch == '\n' {
				count++
			}	
		}
		fmt.Println(count, fileName)
	case "-w":
		content, err := os.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)	
		}
		count := 0
		inWord := false
		for _, ch := range content {
			if ch != ' ' && ch != '\n' && ch != '\t' && ch != '\r' {
				if inWord {
					continue
				} else{
				inWord = true
				count++;
				}
			} else {
				inWord = false
			}
		}
		fmt.Println(count, fileName)
	}

}
