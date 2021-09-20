package piscine

func IsNumeric(s string) bool {
	for _, v := range s {
		if v >= 48 && v <= 57 {
			return true
		}
	}
	return false
}

func Any(f func(string) bool, a []string) bool {
	for _, v := range a {
		if f(v) {
			return true
		}
	}

	return false
}
