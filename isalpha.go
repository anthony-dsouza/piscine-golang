package piscine

func IsAlpha(s string) bool {
	for _, v := range s {
		if v == ' ' {
			return false
		}
	}
	return true
}
