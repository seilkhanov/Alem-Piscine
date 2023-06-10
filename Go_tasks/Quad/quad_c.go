package piscine

import "github.com/01-edu/z01"

func QuadC(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 0; i < y; i++ {
		if i == 0 || i == y-1 {
			for j := 0; j < x; j++ {
				isFirstOrSecond := i == 0 && (j == 0 || j == x-1)
				isThridOrFourth := i == y-1 && (j == 0 || j == x-1)
				if isFirstOrSecond {
					z01.PrintRune('A')
				} else if isThridOrFourth {
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
