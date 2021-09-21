package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		file := os.Stdout
		stat, _ := file.Stat()
		size := stat.Size()
		arr := make([]byte, size)
		file.Read(arr)

		for _, v := range arr {
			z01.PrintRune(rune(v))
		}
	} else {
		for i := range os.Args[1:] {
			file, err := os.Open(os.Args[i+1])
			if err != nil {
				s := "Error: " + err.Error()
				for _, v := range s {
					z01.PrintRune(v)
				}
				return
			}
			stat, _ := file.Stat()
			size := stat.Size()
			arr := make([]byte, size)
			file.Read(arr)

			for _, v := range arr {
				z01.PrintRune(rune(v))
			}

		}
	}
}
