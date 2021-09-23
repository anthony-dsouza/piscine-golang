package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var str string
	neg := false

	if n < 0 {

		n = n * -1
		neg = true
	}
	if n == 0 {
		z01.PrintRune('0')
	} else {
		for n > 0 {
			str = string(rune(n%10+48)) + str
			n = n / 10
		}
	}

	if neg {
		z01.PrintRune('-')
	}

	for _, v := range str {
		z01.PrintRune(v)
	}
}
