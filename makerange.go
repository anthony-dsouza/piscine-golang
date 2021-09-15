package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return make([]int, 0)
	}
	int := make([]int, max)
	slice := int[min:max]

	return slice
}
