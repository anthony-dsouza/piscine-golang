package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	file := os.Args[0][2:]

	for _, v := range file {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
