package piscine

import "github.com/01-edu/z01"

func PrintReverseAlphabet() {
	for c := 'z'; c >= 'a'; c-- {
		err := z01.PrintRune(c)
		if err != nil {
			panic(err)
		}
	}
	z01.PrintRune('\n')
}
