package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	a := 0
	b := 0
	c := 0

	for a <= 9 {
		if c == 7 {
			z01.PrintRune('\n')
			a = 10
		} else {
			if a == 9 {
				if b == 8 && c < 7{
					c++
					b = c + 1
					a = b + 1
				} else if b < 8{
					b++
					a = b + 1
				}
			} else {
				a++
			}
			// comma
			z01.PrintRune(44)
			// space
			z01.PrintRune(32)
			z01.PrintRune(rune(c) + 48)
			z01.PrintRune(rune(b) + 48)
			z01.PrintRune(rune(a) + 48)
		}
	}




}