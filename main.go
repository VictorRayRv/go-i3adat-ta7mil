package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args[1:]) != 2 {
		fmt.Println("Usage: go run . <inputFile> <outputFile>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var finalLines [][]string
	scanner := bufio.NewScanner(file)

	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		
		line = fixingPunc(line)
		line = fixingQuotes(line)
		commands, words, combined := separating(line)
		index := index(combined)
		words = fixingA(words)
		applyTransformation(words,commands,lineNumber,index)

		finalLines = append(finalLines, words)
	}
	filee, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer filee.Close()

	for _,line := range finalLines {
	_, err = filee.WriteString(strings.Join(line," "))
	if err != nil {
		panic(err)
	} 
	_,err = filee.WriteString("\n")
	if err != nil {
		panic(err)
	}
	}
	fmt.Println("File created & written successfully!")
}
