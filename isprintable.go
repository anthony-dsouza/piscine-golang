package piscine

func IsPrintable(s string) bool {
	for _, v := range s {
		if v < 20 {
			return false
		}
	}
	return true
}
