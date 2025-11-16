package main

import ("strings"
		"fmt")

func applyTransformation(words []string,lineNumber int) {
	var err error
	for i := 0 ; i < len(words) ; i++ {
		col := i + 1
		switch {
			case words[i] == "(hex)":
				words[i-1], err = hexToDecimal(words[i-1])
				if err != nil {
					fmt.Printf("[line:%d:column:%d] Error:%v\n",lineNumber,i,err)
				}
			case words[i] == "(bin)":
				words[i-1], err = binToDecimal(words[i-1])
				if err != nil {
					fmt.Printf("[line:%d:column:%d] Error:%v\n",lineNumber,i,err)
				}
			case words[i] == "(up)":
				words[i-1] = toUpper(words[i-1])
			case strings.HasPrefix(words[i], "(up,"):
				nbr, err := extractNumber(words[i])
				if err != nil {
					fmt.Printf("[line:%d:column:%d] Error:%v\n",lineNumber,col,err)
				}
				err = basedOnNbr(words, i, nbr, "up")
				if err != nil {
					fmt.Printf("[line:%d:column:%d] Error:%v\n",lineNumber,col,err)
				}
			case words[i] == "(low)":
				words[i-1] = toLower(words[i-1])
			case strings.HasPrefix(words[i], "(low,"):
				nbr, err := extractNumber(words[i])
				if err != nil {
					fmt.Printf("[line:%d:column:%d] Error:%v\n",lineNumber,col,err)
				}
				err = basedOnNbr(words, i, nbr, "low")
				if err != nil {
					fmt.Printf("[line:%d:column:%d] Error:%v\n",lineNumber,col,err)
				}
			case words[i] == "(cap)":
				words[i-1] = capitalize(words[i-1])
			case strings.HasPrefix(words[i], "(cap,"):
				nbr, err := extractNumber(words[i])
				if err != nil {
					fmt.Printf("[line:%d:column:%d] Error:%v\n",lineNumber,col,err)
				}
				err = basedOnNbr(words, i, nbr, "cap")
				if err != nil {
					fmt.Printf("[line:%d:column:%d] Error:%v\n",lineNumber,col,err)
				}
			}
	}
}
