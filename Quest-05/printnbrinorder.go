package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var aRune []rune

	if n == 0 {
		aRune = append(aRune, 48)
	}

	for i := 0; n > 0; i++ {
		aRune = append(aRune, rune(n%10)+48)
		n = n / 10
	}

	l := len(aRune)
	for j := range aRune {
		if j == l-1 {
		} else {
			for i := 0; i < l-1; i++ {
				if aRune[i] > aRune[i+1] {
					aRune[i], aRune[i+1] = aRune[i+1], aRune[i]
				}
			}
		}
	}

	for _, v := range aRune {
		z01.PrintRune(v)
	}
}
