package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 24 {
		return 0
	} else if nb == 0 {
		return 1
	} else {

		result := nb

		return result * RecursiveFactorial(nb-1)
	}
}
