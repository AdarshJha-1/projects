package main

import (
	"fmt"
	"log"
	"os"
)



func main() {

	// Complicated for me :(
	/*
	byteCountFlag := flag.String("c", "sample.txt", "outputs the number of bytes in a file.")
	lineCountFlag := flag.String("l", "sample.txt", "outputs the numbers of line in a file.")
	flag.Parse()

	content, err := os.ReadFile(*byteCountFlag)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(*lineCountFlag)
	fmt.Print ln(len(content), *byteCountFlag)
	*/

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
			
	}

}
