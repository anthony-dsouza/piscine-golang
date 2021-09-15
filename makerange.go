package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return make([]int, 0)
	}

	size := max - min
	int := make([]int, size)

	for i := 0; i < size; i++ {
		int[i] = min + i
	}

	return int
}
