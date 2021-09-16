package piscine

func SplitWhiteSpaces(s string) []string {
	var str []string

	aRune := []rune(s)
	l := len(aRune) - 1

	var word string

	for i, v := range s {
		if v == 32 || v == 15 || v == 10 {
			if aRune[i+1] == 32 || aRune[i+1] == 15 || aRune[i+1] == 10 {
			} else {
				str = append(str, word)
				word = ""
			}
		} else if i == l {
			word = word + string(v)
			str = append(str, word)
		} else {
			word = word + string(v)
		}
	}

	return str
}
