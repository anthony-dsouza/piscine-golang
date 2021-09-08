package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	for a := 0; a < 10; a++ {
		for b := 0; b < 10; b++ {
			for c := 0; c < 10; c++ {
				if a == 7 && b == 8 && c == 9 {
					z01.PrintRune(rune(a) + 48)
					z01.PrintRune(rune(b) + 48)
					z01.PrintRune(rune(c) + 48)
					z01.PrintRune('\n')
				} else if a < b && b < c {
					z01.PrintRune(rune(a) + 48)
					z01.PrintRune(rune(b) + 48)
					z01.PrintRune(rune(c) + 48)
					z01.PrintRune(44)
					z01.PrintRune(32)
				}
			}
		}
	}
}
