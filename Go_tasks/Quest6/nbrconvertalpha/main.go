package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	path := os.Args[1:]
	upper := false

	if len(path) > 0 && path[0] == "--upper" {
		upper = true
		path = path[1:]
	}

	if len(path) > 0 {
		for _, arg := range path {
			n := atoi(arg)
			if n < 1 || n > 26 {
				z01.PrintRune(' ')
			} else if upper {
				z01.PrintRune(rune('A' - 1 + n))
			} else {
				z01.PrintRune(rune('a' - 1 + n))
			}
		}
		z01.PrintRune('\n')
	}
}

func atoi(s string) int {
	value := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return -1
		}
		value = value*10 + int(r-'0')
	}
	return value
}
