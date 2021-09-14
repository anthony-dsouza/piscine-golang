package piscine

func IsNumeric(s string) bool {
	for _, v := range s {
		if v > '9' || v < '0' {
			return false
		}
	}
	return true
}
