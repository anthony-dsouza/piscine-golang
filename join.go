package piscine

func Join(strs []string, sep string) string {
	var str string
	l := len(strs) - 1
	for i, v := range strs {
		if l == i {
			str = str + v
		} else {
			str = str + v + sep
		}
	}

	return str
}
