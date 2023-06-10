package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		if n == -9223372036854775808 {
			z01.PrintRune('9')
			n = -223372036854775808
		}
		n = -n
	} else if n == 0 {
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
