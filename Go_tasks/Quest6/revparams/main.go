package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[1:]
	for i := len(name) - 1; i >= 0; i-- {
		for _, r := range name[i] {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
