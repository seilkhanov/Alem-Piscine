package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	for i := 0; i < 99; i++ {
		for j := i + 1; j < 100; j++ {
			printCombination(i, j)

			if i != 98 || j != 99 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}

func printCombination(i, j int) {
	printDigits(i)
	z01.PrintRune(' ')
	printDigits(j)
}

func printDigits(num int) {
	z01.PrintRune(rune('0' + num/10))
	z01.PrintRune(rune('0' + num%10))
}

func main() {
	PrintComb2()
}
