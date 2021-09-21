package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var args []string
	for i, v := range os.Args {
		if v == "|" {
			args = os.Args[:i]
			for i := range os.Args[i:] {
				os.Args[i] = ""
			}
		} else {
			args = os.Args
		}
	}
	if len(args) == 1 {
		h := "Hello"
		for _, v := range h {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
		for _, v := range h {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')

		return
	}

	file1, err := os.Open(os.Args[1])
	if err != nil {
		s1 := "ERROR: " + err.Error()

		for _, v := range s1 {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
		exit := "exit status 1"
		for _, v := range exit {
			z01.PrintRune(v)
		}
		return
	}

	stat1, _ := file1.Stat()
	size1 := stat1.Size()

	arr1 := make([]byte, size1)

	file1.Read(arr1)

	for _, v := range arr1 {
		z01.PrintRune(rune(v))
	}

	if len(args) == 3 {
		file2, err := os.Open(os.Args[2])
		if err != nil {
			s2 := "ERROR: " + err.Error()

			for _, v := range s2 {
				z01.PrintRune(v)
			}
			z01.PrintRune('\n')
			return
		}

		stat2, _ := file2.Stat()
		size2 := stat2.Size()

		arr2 := make([]byte, size2)
		file2.Read(arr2)
		for _, v := range arr2 {
			z01.PrintRune(rune(v))
		}
	}
}
