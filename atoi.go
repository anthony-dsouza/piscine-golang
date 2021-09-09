package piscine

func Atoi(s string) int {
	var b int

	neg := false

	for index, i := range s {
		if int(i) == 45 && index == 0 {
			neg = true
		}
	}

	for index, a := range s {
		if int(a) == 45 && index == 0 {
		} else if int(a) == 43 && index == 0 {
		} else if int(a) < 48 || int(a) > 57 {
			return 0
		} else {
			c := int(a) - 48
			b = (b * 10) + c
		}
	}

	if neg {
		c := b * 2
		d := b - c

		return d

	} else {
		return b
	}
}
