package piscine

func Sqrt(nb int) int {
	if nb == 1 {
		return 1
	}
	for i := 0; i < nb; i++ {
		if nb%i > 0 {
		} else if i == 1 {
		} else if i*i == nb {
			return i
		}
	}

	return 0
}
