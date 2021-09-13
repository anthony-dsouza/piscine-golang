package piscine

func NRune(s string, n int) rune {
	aRune := []rune(s)
	len := len(s)
	if n > len || n <= 0 {
		return 0
	}

	return aRune[n-1]
}
