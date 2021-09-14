package piscine

func ToLower(s string) string {
	Rune := []rune(s)
	for index, v := range Rune {
		for i := 'A'; i <= 'Z'; i++ {
			if v == i {
				Rune[index] = v + 32
			}
		}
	}

	return string(Rune)
}
