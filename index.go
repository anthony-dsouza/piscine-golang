package piscine

func Index(s string, toFind string) int {
	find := []rune(toFind)
	length := len(find)

	for i, v := range s {
		if v == find[0] {
			if s[i:length+i] == toFind {
				return i
			}
		}
	}
	return -1
}
