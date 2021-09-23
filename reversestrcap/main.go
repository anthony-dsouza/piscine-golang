package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	var str string

	for _, v := range os.Args[1:] {
		for _, char := range v {
			if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
				if char >= 'A' && char <= 'Z' {
					str = str + string(char+32)
				} else {
					str = str + string(char)
				}
			} else if char >= '0' && char <= '9' {
			} else {
				str = str + string(char)
			}
		}
	}

	aRune := []rune(str)

	l := len(aRune) - 1

	for i, v := range str {
		if v == ' ' {
			if str[i-1] >= 'a' && str[i-1] <= 'z' {
				aRune[i-1] = aRune[i-1] - 32
			}
		} else if i == l && (str[i-1] >= 'a' && str[i-1] <= 'z') {
			aRune[i] = aRune[i] - 32
		}
	}

	fmt.Print(string(aRune))
}
