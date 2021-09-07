package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	a := 0
	b := 0
	c := 0

	z01.PrintRune(rune(a) + 48)
	z01.PrintRune(rune(b) + 48)
	z01.PrintRune(rune(c) + 48)

	for c <= 9 {

		if c == 10 {
			c = 0
		} else if a == 7 && b == 8 && c == 9 {
			z01.PrintRune('\n')
		}

		// comma
		z01.PrintRune(44)
		// space
		z01.PrintRune(32)
		z01.PrintRune(rune(a) + 48)
		z01.PrintRune(rune(b) + 48)
		z01.PrintRune(rune(c) + 48)

		for b <= 8 {
			b++
			if b == 8 {
				b = 0
			}
			// comma
			z01.PrintRune(44)
			// space
			z01.PrintRune(32)
			z01.PrintRune(rune(a) + 48)
			z01.PrintRune(rune(b) + 48)
			z01.PrintRune(rune(c) + 48)
			for a <= 7 {
				a++
				// comma
				z01.PrintRune(44)
				// space
				z01.PrintRune(32)
				z01.PrintRune(rune(a) + 48)
				z01.PrintRune(rune(b) + 48)
				z01.PrintRune(rune(c) + 48)
			}
		}
	}
}
