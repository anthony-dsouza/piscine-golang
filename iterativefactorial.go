package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else {
		result := nb
		for i := 1; i < nb; i++ {
			result = result * (nb - i)
		}

		return result
	}
}
