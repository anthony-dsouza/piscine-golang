package piscine

func Sign(a, b int) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}

func IsSorted(f func(a, b int) int, a []int) bool {
	l := len(a) - 1
	for i := range a {
		if i < l-1 {
			if f(a[i], a[i+1]) == 0 {
				continue
			} else if f(a[i], a[i+1]) > 0 {
				if a[i]-a[i+1] < 0 {
					return false
				} else {
					continue
				}
			} else if f(a[i], a[i+1]) < 0 {
				if a[i]-a[i+1] > 0 {
					continue
				} else {
					return false
				}
			}
		}
	}
	return true
}
