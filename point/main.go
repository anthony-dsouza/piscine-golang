package main

import "github.com/01-edu/z01"

type point struct {
	x string
	y string
}

func main() {
	points := &point{"42", "21"}

	str := "x = " + points.x + ", y = " + points.y

	for _, v := range str {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
