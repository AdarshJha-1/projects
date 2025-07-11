package main

import (
	"fmt"
	"log"
	"os"
	"unicode/utf8"
)


func getNumberOfBytes(fileName string) int {
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)	
	}
	return len(content)
}

func getNumberOfLines(fileName string) int {
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
	return count
}

func getNumberOfWords(fileName string) int {
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
	return count
}
func getNumberOfCharacters(fileName string) int {
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)	
	}
	count := 0
	for len(content) > 0 {
		_, s := utf8.DecodeRune(content)
		content = content[s:]
		count++
	}
	return count
}

func main() {

	if(len(os.Args) < 2) {
		log.Fatal("not enough arguments")
	}

	flag := os.Args[1]
	var fileName string
	if(len(os.Args) == 2 ) {
		fileName = os.Args[1]
	}else {
		fileName = os.Args[2]
	}

	switch flag{
	case "-c": 
		bytes := getNumberOfBytes(fileName)
		fmt.Println(bytes, fileName)
	case "-l":
		lines := getNumberOfLines(fileName)
		fmt.Println(lines, fileName)
	case "-w":
		words := getNumberOfWords(fileName)
		fmt.Println(words, fileName)
	case "-m":
		characters := getNumberOfCharacters(fileName)
		fmt.Println(characters, fileName)
	default:
		lines := getNumberOfLines(fileName)
		words := getNumberOfWords(fileName)
		bytes := getNumberOfBytes(fileName)
		fmt.Println(lines, words, bytes, fileName)
	}
}
