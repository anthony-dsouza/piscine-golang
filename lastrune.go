package piscine

func LastRune(s string) rune {
	aRune := []rune(s)
	len := 0
	for range s {
		len++
	}
	return aRune[len-1]
}
