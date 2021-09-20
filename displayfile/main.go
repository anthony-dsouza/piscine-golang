package main

import (
	"log"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 2 {
		str := "Too many arguments"
		for _, v := range str {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
		return
	}
	if len(os.Args) < 2 {
		str := "File name missing"
		for _, v := range str {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
		return
	}
	file := os.Args[1]

	name, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	stat, err := name.Stat()
	if err != nil {
		log.Fatal(err)
	}

	fsize := stat.Size()

	str := make([]byte, fsize)

	name.Read(str)

	for _, v := range str {
		z01.PrintRune(rune(v))
	}

	name.Close()
}
