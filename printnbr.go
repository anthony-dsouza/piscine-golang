package piscine

import "github.com/01-edu/z01"

func PrintNbr(nbr int) {
	aRune := rune(nbr + 48)
	z01.PrintRune(aRune)
}

func Println() {
	z01.PrintRune('\n')
}
