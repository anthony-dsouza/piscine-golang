package piscine

func ConcatParams(args []string) string {
	var str string
	l := len(args) - 1
	for i, v := range args {
		if i == l {
			str = str + v
		} else {
			str = str + v + "\n"
		}
	}

	return str
}
