package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var lower []rune
	var upper []rune

	for i := 'a'; i <= 'z'; i++ {
		lower = append(lower, i)
	}

	for i := 'A'; i <= 'Z'; i++ {
		upper = append(upper, i)
	}

	if os.Args[1] == "--upper" {
		for _, v := range os.Args[2:] {
			value := 0
			for _, a := range v {
				c := int(a) - 48
				value = (value * 10) + c
			}
			if value > 26 || value < 1 {
				z01.PrintRune(32)
			} else {
				z01.PrintRune(upper[value-1])
			}

		}
	} else {
		for _, v := range os.Args[1:] {
			value := 0
			for _, a := range v {
				c := int(a) - 48
				value = (value * 10) + c
			}
			if value > 26 || value < 1 {
				z01.PrintRune(32)
			} else {
				z01.PrintRune(lower[value-1])
			}
		}
	}
	z01.PrintRune('\n')
}
