package piscine

func Rot14(s string) string {
	var str []rune

	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			v = v + 14
			if v > 'z' {
				v = v - 26
				str = append(str, v)
			} else {
				str = append(str, v)
			}
		} else if v >= 'A' && v <= 'Z' {
			v = v + 14
			if v > 'Z' {
				v = v - 26
				str = append(str, v)
			} else {
				str = append(str, v)
			}
		} else {
			str = append(str, v)
		}
	}
	return string(str)
}
