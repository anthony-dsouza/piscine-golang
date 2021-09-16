package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var str string
	for _, v := range os.Args[1:] {
		str = str + v
	}
	aRune := []rune(str)

	l := len(aRune)
	for j := 0; j < l; j++ {
		for i := range str {
			if i+1 < l {
				if aRune[i] > aRune[i+1] {
					aRune[i], aRune[i+1] = aRune[i+1], aRune[i]
				}
			}
		}
	}

	for _, v := range aRune {
		z01.PrintRune(v)
		z01.PrintRune('\n')
	}
}
