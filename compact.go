package piscine

func Compact(ptr *[]string) int {
	var str []string

	for _, v := range *ptr {
		if v != "" {
			str = append(str, v)
		}
	}

	*ptr = str

	return len(str)
}
