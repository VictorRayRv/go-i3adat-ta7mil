package main

import "strings"

func fixingPunc(text string) string {
	result := ""
	runes := []rune(text)

	i := 0 
	prevWasSpace := false 

	for i < len(runes) {
		char := runes[i]

		if char == ' ' {
			if !prevWasSpace {
				result += " "
				prevWasSpace = true
			}
			i++
			continue
		}

		prevWasSpace = false 

		if char == '.' || char == ',' || char == '!' || char == '?' || char == ';' || char == ':' {
			if len(result) > 0 && result[len(result)-1] == ' ' {
				result = result[:len(result)-1]
			}
			
			for i < len(runes) && (runes[i] == '.' || runes[i] == ',' || runes[i] == '!' || runes[i] == '?' || runes[i] == ';' || runes[i] == ':') {
				result += string(runes[i])
				i++
			}
			if i < len(runes) && ((runes[i] >= 'A' && runes[i] <= 'Z') || (runes[i] >= 'a' && runes[i] <= 'z')) {
				result += " "
				prevWasSpace = true
			}
			continue
		}

		result += string(char)
		i++
	}
	return result
}

func fixingQuotes(text string) string {
	result := ""
	runes := []rune(text)
	inQuotes := false
	
	for i := 0; i < len(runes); i++ {
		char := runes[i]
		if char == '\'' {
			if !inQuotes {
				inQuotes = true
				result += "'"
				for i+1 < len(runes) && runes[i+1] == ' ' {
					i++
				}
			} else {
				inQuotes = false
				if len(result) > 0 && result[len(result)-1] == ' ' {
					result = result[:len(result)-1]
				}
				result += "'"
			}
			continue
		}
		result += string(char)
	}
	return result
}

func fixingA(words []string) []string {
	if len(words) == 0 {
		return words
	}

	for i := 0 ; i < len(words)-1 ; i++ {
		if toLower(words[i]) == "a" {
			next := toLower(words[i+1])
			if strings.HasPrefix(next,"'") && strings.HasSuffix(next,"'"){
				next = strings.ReplaceAll(next,"'","")
			}
			if startsWithVowelsOrH(next) {
				words[i] = "an"
			}
		}
	}
	return words
}