package piscine

func BasicAtoi(s string) int {
	var b int

	for _, a := range s {
		c := int(a) - 48
		b = (b * 10) + c

	}

	return b
}
