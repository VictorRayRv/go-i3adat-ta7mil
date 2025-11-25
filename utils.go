package main

import (
	"fmt"
	"strconv"
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

func startsWithVowelsOrH(word string) bool {
	if len(word) == 0 {
		return false
	}
	first := rune(word[0])
	return first == 'a' || first == 'e' || first == 'i' || first == 'o' || first == 'u' || first == 'h'
}

func commandParsing(str string) Command {
	for i := 0; i < len(str) ; i++ {
		if str[i] == ' ' {
			continue
		}
		if str[i] == '(' {
			j := i + 1 

			name := ""
			for j < len(str) && str[j] >= 'a' && str[j] <= 'z' {
				name += string(str[j])
				j++
			}

			if name != "cap" && name != "up" && name != "low" && name != "bin" && name != "hex" {
				continue
			}

			value := 1 

			if j < len(str) && str[j] == ',' {
				j++
				numStr := ""

				for j < len(str) && str[j] >= '0' && str[j] <= '9' {
					numStr += string(str[j])
					j++
				}

				if numStr == "" {
					return Command{"",0,fmt.Errorf("missing number after comma")}
				}
				n,err := strconv.Atoi(numStr)
				if err != nil {
					return Command{"",0,fmt.Errorf("invalid number")}
				}
				value = n
			}
			if j >= len(str) || str[j] != ')' {
				return Command{"",0,fmt.Errorf("missing closing parenthesis")}
	
			}
			return Command{name, value, nil}
		}
	}
	return Command{"",0,fmt.Errorf("no command found")}
}

func index(slice []string) int {
	for i, w := range slice {
		cmd := commandParsing(w)
		if cmd.Err == nil {
			return i
		}
	}
	return 0
}