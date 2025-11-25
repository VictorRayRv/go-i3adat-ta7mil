package main

import "fmt"

type Command struct {
	Name string
	Value int
	Err error
}

func separating(line string) ([]Command, []string, []string) {
	combined := []string{}
	words := []string{}
	commands := []Command{}
	runes := []rune(line)
	i := 0
	word := ""

	for i < len(runes) {
		if runes[i] == '(' {
			if word != "" {
				combined = append(combined, word)
				words = append(words, word)
				word = ""
			}

			cmdStr := ""
			for i < len(runes) && runes[i] != ')' {
				cmdStr += string(runes[i])
				i++
			}
			if i < len(runes) && runes[i] == ')' {
				cmdStr += ")"
				i++
			}

			cmd := commandParsing(cmdStr)
			combined = append(combined, cmdStr)
			commands = append(commands, cmd)
			
		} else if runes[i] == ' ' {
			if word != "" {
				combined = append(combined, word)
				words = append(words, word)
				word = ""
			}
			i++
		} else {
			word += string(runes[i])
			i++
		}
	}

	if word != "" {
		combined = append(combined, word)
		words = append(words, word)
	}

	return commands, words , combined
}


func applyTransformation(words []string,commands []Command, lineNumber int, index int) {
	for _,cmd := range commands {
		if cmd.Err != nil {
			continue
		}

		col := index+1

		switch cmd.Name {
		case "cap", "up" , "low" :
			err := basedOnNbr(words, index, cmd.Value, cmd.Name)
			if err != nil {
				fmt.Printf("[line:%d:column:%d] Error:%v\n", lineNumber,col, err)
			}
		case "hex":
			newWord, err := hexToDecimal(words[index-1])
			if err != nil {
				fmt.Printf("[line:%d:column:%d] Error:%v\n", lineNumber,col, err)
			} else {
				words[index-1] = newWord
			}
		case "bin":
			newWord, err := binToDecimal(words[index-1])
			if err != nil {
				fmt.Printf("[line:%d:column:%d] Error:%v\n", lineNumber,col, err)
			} else {
				words[index-1] = newWord
			}
		}
	}
}
