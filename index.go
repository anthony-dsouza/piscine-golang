package piscine

func Index(s string, toFind string) int {
	find := []rune(toFind)
	length := len(find)
	str := []rune(s)
	s_len := len(str)

	if length <= 0 {
		return -1
	}

	for i, v := range s {
		if v == find[0] {
			if length+i > s_len {
			} else {
				if s[i:length+i] == toFind {
					return i
				}
			}
		}
	}
	return -1
}
