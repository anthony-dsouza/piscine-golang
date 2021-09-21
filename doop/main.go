package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	Args := os.Args[1:]
	value1, _ := strconv.Atoi(Args[0])
	value2, _ := strconv.Atoi(Args[2])

	op := Args[1]

	if op == "+" {
		value := value1 + value2
		var str string
		if value < -2147483648 || value > 2147483647 {
			return
		}
		if value < 0 {
			z01.PrintRune('-')
			value = value * (-1)
			c := rune(value%10) + 48
			str = string(c) + str
			value = value / 10
		}
		for value > 0 {
			c := rune(value%10) + 48
			str = string(c) + str
			value = value / 10
		}

		for _, v := range str {
			z01.PrintRune(v)
		}
	}

	if op == "-" {
		value := value1 - value2
		var str string

		if value < -2147483648 || value > 2147483647 {
			return
		}
		if value < 0 {
			z01.PrintRune('-')
			value = value * (-1)
			c := rune(value%10) + 48
			str = string(c) + str
			value = value / 10
		}

		for value > 0 {
			c := rune(value%10) + 48
			str = string(c) + str
			value = value / 10
		}

		for _, v := range str {
			z01.PrintRune(v)
		}

	}

	if op == "*" {
		value := value1 * value2
		var str string
		if value < -2147483648 || value > 2147483647 {
			return
		}
		if value < 0 {
			z01.PrintRune('-')
			value = value * (-1)
			c := rune(value%10) + 48
			str = string(c) + str
			value = value / 10
		}

		for value > 0 {
			c := rune(value%10) + 48
			str = string(c) + str
			value = value / 10
		}

		for _, v := range str {
			z01.PrintRune(v)
		}
	}

	if op == "/" {
		if value2 == 0 {
			str := "No division by 0"
			for _, v := range str {
				z01.PrintRune(v)
			}
		} else {
			value := value1 / value2
			var str string

			if value < -2147483648 || value > 2147483647 {
				return
			}

			for value > 0 {
				c := rune(value%10) + 48
				str = string(c) + str
				value = value / 10
			}

			for _, v := range str {
				z01.PrintRune(v)
			}
		}
	}

	if op == "%" {
		if value2 == 0 {
			str := "No modulo by 0"
			for _, v := range str {
				z01.PrintRune(v)
			}
		} else {
			value := value1 % value2
			var str string
			if value < -2147483648 || value > 2147483647 {
				return
			}

			for value > 0 {
				c := rune(value%10) + 48
				str = string(c) + str
				value = value / 10
			}

			for _, v := range str {
				z01.PrintRune(v)
			}
		}
	}
}
