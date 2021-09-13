package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 24 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		result := nb
		for i := 1; i < nb; i++ {
			result = result * (nb - i)
		}

		return result
	}
}
