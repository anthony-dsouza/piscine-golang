package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	l := len(os.Args[1:])
	for j := 0; j < l; j++ {
		for i := range args {
			if i+1 < l {
				if args[i] > args[i+1] {
					args[i], args[i+1] = args[i+1], args[i]
				}
			}
		}
	}

	for _, v := range os.Args[1:] {
		for _, a := range v {
			z01.PrintRune(a)
		}

		z01.PrintRune('\n')
	}
}
