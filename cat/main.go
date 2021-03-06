package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		file, _ := io.ReadAll(os.Stdin)

		for _, v := range file {
			z01.PrintRune(rune(v))
		}
	} else {
		for i := range os.Args[1:] {
			file, err := os.Open(os.Args[i+1])
			if err != nil {
				s := "ERROR: " + err.Error()
				for _, v := range s {
					z01.PrintRune(v)
				}
				z01.PrintRune('\n')
				os.Exit(1)
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
