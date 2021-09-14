package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune(48)
	}

	for i := 0; n > 0; i++ {
		z01.PrintRune(rune(n%10) + 48)
		n = n / 10
	}
}
