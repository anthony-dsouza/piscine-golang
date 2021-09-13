package piscine

func BasicAtoi2(s string) int {
	var b int

	for _, a := range s {
		if int(a) < 48 || int(a) > 57 {
			return 0
		} else {
			c := int(a) - 48
			b = (b * 10) + c
		}
	}

	return b
}
