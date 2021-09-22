package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	num := 0
	for _, v := range a {
		if v > num {
			num = v
		}
	}
	return num
}
