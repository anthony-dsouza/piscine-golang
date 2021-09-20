package main

import "fmt"

type point struct {
	x int
	y int
}

func main() {
	points := &point{42, 21}

	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
