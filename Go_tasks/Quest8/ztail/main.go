package main

import (
	"fmt"
	"os"
)

func main() {
	fileNames := os.Args
	// index := 0
	index := Atoi(fileNames[2])
	for i := 3; i < len(fileNames); i++ {
		data, err := os.ReadFile(fileNames[i])
		if err != nil {
			if i != 3 {
				fmt.Printf("\n")
			}
			start := len(data) - index
			if index > len(data)-1 {
				fmt.Printf("==> %s <==\n%s", fileNames[i], string(data))
				os.Exit(1)
				fmt.Printf("\n")
			}
			if index == len(data)-1 {
				fmt.Printf("==> %s <==\n%s", fileNames[i], string(data)[1:])
				os.Exit(1)
				fmt.Printf("\n")
			}
			fmt.Printf("==> %s <==\n%s", fileNames[i], string(data)[start:])
		} else {
			fmt.Printf("%v", err)
			fmt.Printf("\n")
			if len(fileNames) == 4 {
				os.Exit(1)
			}
		}
	}
}

func Atoi(s string) int {
	result := 0
	sign := 1
	i := 0

	if len(s) > 0 && s[0] == '-' {
		sign = -1
		i++
	} else if len(s) > 0 && s[0] == '+' {
		i++
	}

	for ; i < len(s); i++ {
		digit := int(s[i] - '0')

		if digit < 0 || digit > 9 {
			result = 0
			break
		} else {
			result = result*10 + digit
		}
	}
	return result * sign
}
