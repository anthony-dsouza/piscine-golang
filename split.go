package piscine

func Split(s, sep string) []string {
	l := len(sep)

	var strings []string

	var str string

	for i, v := range s {
		if len(s)-1 == i {
			str = str + string(v)
			str = str[l-1:]
			strings = append(strings, str)
		} else if v == rune(sep[0]) {
			if i+l > len(s)-1 {
				str = str + string(v)
			} else {
				if s[i:i+l] == sep {
					if str[0] == sep[1] {
						str = str[l-1:]
					}
					strings = append(strings, str)
					str = ""

				} else {
					str = str + string(v)
				}
			}
		} else {
			str = str + string(v)
		}
	}

	return strings
}
