package piscine

func Capitalize(s string) string {
	aRune := []rune(s)

	for i, v := range aRune {
		if v >= 'A' && v <= 'Z' || v >= 'a' && v <= 'z' {
			if v >= 'A' && v <= 'Z' {
				aRune[i] = v + 32
			}
		}
	}

	for i, v := range aRune {
		if v >= 'a' && v <= 'z' {
			if i-1 < 0 {
				aRune[i] = v - 32
			} else if (aRune[i-1] >= 'a' && aRune[i-1] <= 'z') || (aRune[i-1] >= 'A' && aRune[i-1] <= 'Z') || (aRune[i-1] >= '0' && aRune[i-1] <= '9') {
			} else {
				aRune[i] = v - 32
			}
		}
	}
	return string(aRune)
}
