package main

import (
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, err := strconv.Atoi(os.Args[1])
	if a == 9223372036854775807 || a == -9223372036854775807 {
		return
	}
	if err != nil {
		return
	}

	b, err := strconv.Atoi(os.Args[3])
	if b == 9223372036854775807 || b == -9223372036854775807 {
		return
	}
	if err != nil {
		return
	}

	for _, v := range os.Args[2] {
		if v == '+' {
			c := a + b
			d := strconv.Itoa(c)
			e := []byte(d)
			os.Stdout.Write(e)
			f := []byte("\n")
			os.Stdout.Write(f)
		}
		if v == '-' {
			c := a - b
			d := strconv.Itoa(c)
			e := []byte(d)
			os.Stdout.Write(e)
			f := []byte("\n")
			os.Stdout.Write(f)
		}
		if v == '*' {
			c := a * b
			d := strconv.Itoa(c)
			e := []byte(d)
			os.Stdout.Write(e)
			f := []byte("\n")
			os.Stdout.Write(f)
		}
		if v == '/' {
			if os.Args[3] == "0" {
				d := "No division by 0"
				e := []byte(d)
				os.Stdout.Write(e)
				return
			}
			c := a / b
			d := strconv.Itoa(c)
			e := []byte(d)
			os.Stdout.Write(e)
			f := []byte("\n")
			os.Stdout.Write(f)
		}
		if v == '%' {
			if os.Args[3] == "0" {
				d := "No modulo by 0"
				e := []byte(d)
				os.Stdout.Write(e)
				return
			}
			c := a % b
			d := strconv.Itoa(c)
			e := []byte(d)
			os.Stdout.Write(e)
			f := []byte("\n")
			os.Stdout.Write(f)
		}

	}
}
