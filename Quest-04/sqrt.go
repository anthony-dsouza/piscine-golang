package piscine

func Sqrt(nb int) int {
	for i := 1; i < nb; i++ {
		if nb%i > 0 {
		} else if i == 1 {
		} else if i*i == nb {
			return i
		}
	}

	return 0
}
