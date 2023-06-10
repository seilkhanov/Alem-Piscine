package piscine

import (
	"github.com/01-edu/z01"
)

func DescendComb() {
	for i := 99; i >= 1; i-- {
		for j := i - 1; j >= 0; j-- {
			printCombination(i, j)

			if i != 1 || j != 0 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}

func printCombination(a, b int) {
	printDigits(a)
	z01.PrintRune(' ')
	printDigits(b)
}

func printDigits(a int) {
	if a != 0 {
		z01.PrintRune(rune(a/10 + '0'))
		z01.PrintRune(rune(a%10 + '0'))
	} else {
		z01.PrintRune(rune(a + '0'))
		z01.PrintRune(rune(a + '0'))
	}
}
