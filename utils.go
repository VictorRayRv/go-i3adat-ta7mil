package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func hexToDecimal(hex string) (string, error) {
	decimal := 0
	notHex := []rune{}
	for _, r := range hex {
		var value int
		switch {
		case r >= '0' && r <= '9':
			value = int(r - '0')
		case r >= 'A' && r <= 'F':
			value = 10 + int(r -'A')
		case r >= 'a' && r <= 'f':
			value = 10 + int(r -'a')
		default:
			notHex = append(notHex, r)
			continue
		}
		decimal = decimal*16 + value
	}
	if len(notHex) > 0 {
		return hex, fmt.Errorf("invalid hex characters: %q", string(notHex))
	}
	return strconv.Itoa(decimal), nil
}

func binToDecimal(bin string) (string, error) {
	decimal := 0
	notbin := []rune{}
	for _, r := range bin {
		var value int
		switch r {
		case '0':
			value = 0
		case '1':
			value = 1
		default:
			notbin = append(notbin, r)
			continue
		}
		decimal = decimal*2 + value
	}
	if len(notbin) > 0 {
		return bin, fmt.Errorf("invalid binary characters: %q", string(notbin))
	}
	return strconv.Itoa(decimal), nil
}

func toUpper(word string) string {
	result := ""
	for _, char := range word {
		if char >= 'a' && char <= 'z' {
			result = result + string(char-32)
		} else {
			result = result + string(char)
		}
	}
	return result
}

func toLower(word string) string {
	result := ""
	for _, char := range word {
		if char >= 'A' && char <= 'Z' {
			result = result + string(char+32)
		} else {
			result = result + string(char)
		}
	}
	return result
}

func basedOnNbr(words []string, i int, nbr int, keyWord string,) error {
	if nbr > len(words)-1 {
		return fmt.Errorf("invalid count for (%s): asked to modify %d words, but only %d available",
		keyWord, nbr, len(words)-1)
	}
	start := i - nbr
	if start < 0 {
		start = 0
	}
	for j := start; j < i; j++ {
		switch keyWord {
		case "up":
			words[j] = toUpper(words[j])
		case "low":
			words[j] = toLower(words[j])
		case "cap":
			words[j] = capitalize(words[j])
		}
	}
	return nil
}

func extractNumber(word string) (int, error) {
	numStr := ""
	for _, r := range word {
		if r >= '0' && r <= '9' {
			numStr += string(r)
		}
	}
	if numStr == "" {
		return 0, fmt.Errorf("no number found in the instance:%q", word)
	}
	return strconv.Atoi(numStr)
}

func capitalize(word string) string {
	first := word[0]
	if first >= 'a' && first <= 'z' {
		first = first - 32
	}
	rest := ""
	for _, ch := range word[1:] {
		if ch >= 'A' && ch <= 'Z' {
			ch = ch + 32
		}
		rest += string(ch)
	}
	return string(first) + rest
}

func filtring(words []string) string {
	cleanWords := []string{}
	for _, w := range words {
		if strings.HasPrefix(w, "(") && strings.HasSuffix(w, ")") {
			continue
		}
		cleanWords = append(cleanWords, w)
	}
	return strings.Join(cleanWords, " ")
}

func fixingPunc(text string) string {
	result := ""
	runes := []rune(text)

	for i := 0; i < len(runes); i++ {
		char := runes[i]

		if strings.ContainsRune(".,!?;:", char) {

			if len(result) > 0 && result[len(result)-1] == ' ' {
				result = result[:len(result)-1]
			}

			result += string(char)

			if i+1 < len(runes) && unicode.IsLetter(runes[i+1]) {
				result += " "
			}
		} else {
			result += string(char)
		}
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

func startsWithVowelsOrH(word string) bool {
	if len(word) == 0 {
		return false
	}
	first := rune(word[0])
	return first == 'a' || first == 'e' || first == 'i' || first == 'o' || first == 'u' || first == 'h'
}