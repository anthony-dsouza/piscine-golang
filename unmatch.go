package piscine

func Unmatch(a []int) int {
	count := 0
	var num int

	for _, v := range a {
		for _, c := range a {
			if v == c {
				count++
			}
		}
		if count != 2 {
			num = v
			return num
		}
		count = 0
	}
	return -1
}
