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

	var outputLines []string
	scanner := bufio.NewScanner(file)

	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		words := strings.Fields(line)

		for i := 0; i < len(words); i++ {
			if strings.HasPrefix(words[i], "(") && strings.HasSuffix(words[i], ",") && i+1 < len(words) {
				words[i] = words[i] + words[i+1]
				words = append(words[:i+1], words[i+2:]...)
			}
			if strings.HasPrefix(words[i], "(") && strings.Contains(words[i], ",") {
				words[i] = strings.ReplaceAll(words[i], " ", "")
			}
		}
		applyTransformation(words,lineNumber)

		lineS := strings.Join(words, " ")
		newLineS := fixingPunc(lineS)
		newLineS = fixingQuotes(newLineS) 
		wordsSlice := strings.Fields(newLineS)
		wordsSlice = fixingA(wordsSlice)
		newLineS = filtring(wordsSlice)
		outputLines = append(outputLines, newLineS)
	}
	err = os.WriteFile(outputFile, []byte(strings.Join(outputLines, "\n")), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}
