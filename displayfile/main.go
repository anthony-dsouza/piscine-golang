package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Print("Too many arguments")
		return
	}
	if len(os.Args) < 2 {
		fmt.Print("File name missing")
		return
	}
	file := os.Args[1]

	name, err := os.Open(file)
	if err != nil {
		fmt.Print(err.Error())
	}

	stat, err := name.Stat()
	if err != nil {
		fmt.Print(err.Error())
	}

	fsize := stat.Size()

	str := make([]byte, fsize)

	name.Read(str)

	fmt.Print()

	name.Close()
}
