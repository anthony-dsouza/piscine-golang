package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func Even(n int) int {
	if n == 0 {
		return 0
	}
	if n%2 == 0 {
		return 1
	}
	return 0
}

func isEven(nbr int) bool {
	if Even(nbr) == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	lengthOfArg := len(os.Args[1:])
	EvenMsg := "I have an number even of arguments"
	OddMsg := "I have an odd number of arguments"
	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
