package piscine

func Join(strs []string, sep string) string {
	var str string

	for _, v := range strs {
		str = str + v + ":"
	}

	return str
}
