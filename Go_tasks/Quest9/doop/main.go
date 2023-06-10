package main

import (
	"os"

	"github.com/01-edu/z01"
)

func calculate(a, b int, operator string) *int {
	switch operator {
	case "+":
		return add(a, b)
	case "-":
		return subtract(a, b)
	case "*":
		return multiply(a, b)
	case "/":
		return divide(a, b)
	case "%":
		return mod(a, b)
	default:
		return nil
	}
}

func add(a, b int) *int {
	result := a + b
	return &result
}

func subtract(a, b int) *int {
	result := a - b
	return &result
}

func multiply(a, b int) *int {
	result := a * b
	return &result
}

func divide(a, b int) *int {
	if b == 0 {
		printDivError()
		return nil
	}

	result := a / b
	return &result
}

func mod(a, b int) *int {
	if b == 0 {
		printModError()
		return nil
	}

	result := a % b
	return &result
}

func main() {
	setup := os.Args[1:]

	if len(setup) != 3 {
		return
	}

	valid := IsNumeric(setup[0]) && IsNumeric(setup[2])

	if valid {
		a := Atoi(setup[0])
		b := Atoi(setup[2])
		operator := setup[1]

		result := calculate(a, b, operator)

		if result != nil {
			if *result < 2147483647 && *result > -2147483648 {
				PrintNbr(*result)
				z01.PrintRune('\n')
			}
		}
	}
}

func Atoi(s string) int {
	result := 0
	sign := 1
	i := 0

	if len(s) > 0 && s[0] == '-' {
		sign = -1
		i++
	} else if len(s) > 0 && s[0] == '+' {
		i++
	}

	for ; i < len(s); i++ {
		digit := int(s[i] - '0')

		if digit < 0 || digit > 9 {
			result = 0
			break
		} else {
			result = result*10 + digit
		}
	}
	return result * sign
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n == 0 {
		z01.PrintRune('0')
	}
	PrintDigit(n)
}

func PrintDigit(n int) {
	if n > 0 {
		PrintDigit(n / 10)
		z01.PrintRune(rune(n%10 + '0'))
	}
}

func IsNumeric(s string) bool {
	for i := range s {
		if (s[i] >= 48 && s[i] <= 57) || s[i] == '-' {
			return true
		}
	}
	return false
}

func printDivError() {
	err := "No division by 0"
	for _, letter := range err {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}

func printModError() {
	err := "No modulo by 0"
	for _, letter := range err {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}
