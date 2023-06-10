package piscine

import "github.com/01-edu/z01"

func QuadD(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 0; i < y; i++ {
		if i == 0 || i == y-1 {
			for j := 0; j < x; j++ {
				isFirstOrThird := (i == 0 || i == y-1) && j == 0
				isSecondOrFourth := (i == 0 || i == y-1) && j == x-1
				if isFirstOrThird {
					z01.PrintRune('A')
				} else if isSecondOrFourth {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('B')
				}
			}
			z01.PrintRune('\n')
		} else {
			for j := 0; j < x; j++ {
				if j == 0 || j == x-1 {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
