package piscine

import (
	"sort"
)

func MaxWordCountN(text string, n int) map[string]int {
	words := make(map[string]int)

	var str string

	for _, v := range text {
		if v == ' ' {
			value, ok := words[str]
			if ok {
				words[str] = value + 1
			} else {
				words[str] = 1
			}
			str = ""
		} else {
			str = str + string(v)
		}
	}

	values := make([]string, 0, len(words))
	for k := range words {
		values = append(values, k)
	}

	sort.Slice(values, func(i, j int) bool {
		return words[values[i]] < words[values[j]]
	})

	twords := make(map[string]int)

	for i := len(words) - 1; i >= len(words)-n; i-- {
		twords[values[i]] = words[values[i]]
	}

	return twords
}
