package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	a := []byte(s)

	// b := len(a)

	// for i := 0; i < b; i++ {
	// 	j := rune(a[i])
	// 	z01.PrintRune(j)
	// }

	for _, letter := range a {
		z01.PrintRune(rune(letter))
	}
}
