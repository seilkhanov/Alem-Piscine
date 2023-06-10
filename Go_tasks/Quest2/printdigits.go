package piscine

import "github.com/01-edu/z01"

func PrintDigits() {
	for d := '0'; d <= '9'; d++ {
		err := z01.PrintRune(d)
		if err != nil {
			panic(err)
		}
	}
	z01.PrintRune('\n')
}
