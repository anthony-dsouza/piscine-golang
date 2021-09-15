package piscine

func MakeRange(min, max int) []int {
	if min >= max || min < 0 || max < 0 {
		return make([]int, 0)
	}
	int := make([]int, max)
	slice := int[min:max]

	return slice
}
