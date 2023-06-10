package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	path := os.Args[1:]

	for i := 0; i < len(path)-1; i++ {
		for j := 0; j < len(path)-1-i; j++ {
			if path[j] > path[j+1] {
				path[j], path[j+1] = path[j+1], path[j]
			}
		}
	}

	for _, arg := range path {
		for _, r := range arg {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
