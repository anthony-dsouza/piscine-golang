package main

import "github.com/01-edu/z01"

func main() {
	var i rune = 97

	for i <= 122 {
		z01.PrintRune(i)
		i++
	}

	if i > 122 {
		z01.PrintRune('\n')
	}
}
