package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for a := 0; a < 10; a++ {
		for b := 0; b < 10; b++ {
			for c := 0; c < 10; c++ {
				for d := 0; d < 10; d++ {
					i := (a * 10) + b
					j := (c * 10) + d
					if i == 98 && j == 99 {
						z01.PrintRune(rune(a) + 48)
						z01.PrintRune(rune(b) + 48)
						z01.PrintRune(32)
						z01.PrintRune(rune(c) + 48)
						z01.PrintRune(rune(d) + 48)
						z01.PrintRune('\n')
					} else if j > i {
						z01.PrintRune(rune(a) + 48)
						z01.PrintRune(rune(b) + 48)
						z01.PrintRune(32)
						z01.PrintRune(rune(c) + 48)
						z01.PrintRune(rune(d) + 48)
						z01.PrintRune(44)
						z01.PrintRune(32)
					}

				}
			}
		}
	}
}
