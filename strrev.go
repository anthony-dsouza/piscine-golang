package piscine

func StrRev(s string) string {
	a := []byte(s)

	count := len(a)

	for i, j := 0, count-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	return string(a)
}
