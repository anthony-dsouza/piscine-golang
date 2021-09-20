package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
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

	fmt.Print(string(str))

	name.Close()
}
