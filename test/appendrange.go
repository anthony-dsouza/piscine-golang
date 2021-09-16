package piscine

func AppendRange(min, max int) []int {
	var int []int

	if min >= max {
		return int
	}

	for i := min; i < max; i++ {
		int = append(int, i)
	}

	return int
}
