package piscine

func MaxWordCountN(text string, n int) map[string]int {
	words := make(map[string]int)

	var str string

	for _, v := range text {
		if v == ' ' {
			value, ok := words[str]
			if ok {
				words[str] = value + 1
			} else {
				words[str] = 1
			}
			str = ""
		} else {
			str = str + string(v)
		}
	}

	var values []int
	for k := range words {
		values = append(values, words[k])
	}

	return words
}
