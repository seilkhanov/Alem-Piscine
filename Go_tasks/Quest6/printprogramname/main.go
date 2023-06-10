package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	path := os.Args[0]
	baseName := getBaseName(path)
	for _, char := range baseName {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func getBaseName(path string) string {
	lastSlash := -1
	for i, char := range path {
		if char == '/' {
			lastSlash = i
		}
	}

	if lastSlash >= 0 && lastSlash < len(path)-1 {
		return path[lastSlash+1:]
	} else {
		return path
	}
}
