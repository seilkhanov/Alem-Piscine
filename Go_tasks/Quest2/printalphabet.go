package piscine

import "github.com/01-edu/z01"

func PrintAlphabet() {
	for c := 'a'; c <= 'z'; c++ {
		err := z01.PrintRune(c)
		if err != nil {
			panic(err)
		}
	}
	z01.PrintRune('\n')
}
