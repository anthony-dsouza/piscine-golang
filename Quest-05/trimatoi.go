package piscine

func TrimAtoi(s string) int {
	var num int
	neg := false

	for _, v := range s {
		if v >= '0' && v <= '9' {
			num = (num * 10) + (int(v) - 48)
		}
	}
	var n int
	for i, v := range s {
		if v == '-' {
			n = i
		}
	}

	for _, v := range s[:n] {
		if v >= '0' && v <= '9' {
			neg = false
		} else {
			neg = true
		}
	}

	if neg {
		num = num * -1
	}

	return num
}
