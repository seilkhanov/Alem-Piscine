package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	slice := []int{}

	if n == 0 {
		z01.PrintRune('0')
	}

	for n > 0 {
		slice = append(slice, n%10)
		n /= 10
	}

	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

	for i := 0; i < len(slice); i++ {
		z01.PrintRune(rune(slice[i]) + '0')
	}
}
