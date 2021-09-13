package piscine

func ToUpper(s string) string {
	Rune := []rune(s)
	for index, v := range Rune {
		for i := 'a'; i <= 'z'; i++ {
			if v == i {
				Rune[index] = v - 32
			}
		}
	}

	return string(Rune)
}
