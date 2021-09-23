package piscine

import (
	"sort"
)

func MaxWordCountN(text string, n int) map[string]int {
	words := make(map[string]int)

	var str string
	l := len(text) - 1
	for i, v := range text {
		if v == ' ' || i == l {
			if i == l && v == 32 {
				words[str]++
				str = ""
				words[str]++
			} else if i == l {
				str = str + string(v)
				_, ok := words[str]
				if ok {
					words[str]++
				} else {
					words[str] = 1
				}
			} else {
				_, ok := words[str]
				if ok {
					words[str]++
					str = ""
				} else {
					words[str] = 1
					str = ""
				}

			}
		} else {
			str = str + string(v)
		}
	}

	values := make([]string, 0, len(words))
	for k := range words {
		values = append(values, k)
	}

	sort.Slice(values, func(i, j int) bool {
		if words[values[i]] > words[values[j]] {
			return true
		}
		if words[values[i]] < words[values[j]] {
			return false
		}
		return values[i] < values[j]
	})

	twords := make(map[string]int)

	for i := 0; i < n; i++ {
		twords[values[i]] = words[values[i]]
	}

	return twords
}
