package piscine

func SortIntegerTable(table []int) {
	count := len(table)

	for i, j := 0, count-1; i < j; i, j = i+1, j-1 {
		table[i], table[j] = table[j], table[i]
	}
}
