package piscine

func IsPrime(n int) bool {
	if n == 1 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	} else {
		return true
	}
}

func Map(f func(int) bool, a []int) []bool {
	var b []bool

	for _, v := range a {
		if f(v) {
			b = append(b, true)
		} else {
			b = append(b, false)
		}
	}

	return b
}
