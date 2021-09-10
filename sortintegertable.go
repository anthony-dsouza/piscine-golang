package piscine

func SortIntegerTable(table []int) {
	count := len(table)
	for a := 0; a < count; a++ {
		for i := range table {
			if i == count-1 {
			} else {
				if table[i] > table[i+1] {
					j := table[i+1]
					table[i+1] = table[i]
					table[i] = j
				}
			}
		}
	}
}
