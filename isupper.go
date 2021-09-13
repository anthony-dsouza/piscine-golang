package piscine

func IsUpper(s string) bool {
	for _, v := range s {
		for i := 'a'; i <= 'z'; i++ {
			if i == v {
				return false
			}
		}
	}

	return true
}
