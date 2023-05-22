package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	a := rectangle{'o', 'o', 'o', 'o', '-', '|'}
	QuadHelper(x, y, a)
}

func QuadB(x, y int) {
	b := rectangle{'/', '\\', '/', '\\', '*', '*'}
	QuadHelper(x, y, b)
}

func QuadC(x, y int) {
	c := rectangle{'A', 'A', 'C', 'C', 'B', 'B'}
	QuadHelper(x, y, c)
}

func QuadD(x, y int) {
	d := rectangle{'A', 'C', 'C', 'A', 'B', 'B'}
	QuadHelper(x, y, d)
}

func QuadE(x, y int) {
	e := rectangle{'A', 'C', 'A', 'C', 'B', 'B'}
	QuadHelper(x, y, e)
}

func QuadHelper(x, y int, r rectangle) {
	if x < 1 || y < 1 {
		return
	}

	if x == 1 && y == 1 {
		z01.PrintRune(r._1)
		z01.PrintRune('\n')
	} else if x == 1 {
		z01.PrintRune(r._1)
		for i := y - 2; i > 0; i-- {
			z01.PrintRune(r.side2)
		}
		z01.PrintRune(r._4)
	} else if y == 1 {
		z01.PrintRune(r._1)
		z01.PrintRune('\n')
		for i := x - 2; i > 0; i-- {
			z01.PrintRune(r.side1)
		}
		z01.PrintRune(r._2)

	} else {
		z01.PrintRune(r._1)

		for i := x - 2; i > 0; i-- {
			z01.PrintRune(r.side1)
		}
		z01.PrintRune(r._2)
		z01.PrintRune('\n')

		for i := x - 2; i > 0; i-- {
			z01.PrintRune(r.side2)
			for j := x - 2; j > 0; j-- {
				z01.PrintRune(' ')
			}
			z01.PrintRune(r.side2)
			z01.PrintRune('\n')

		}

		z01.PrintRune(r._3)

		for i := x - 2; i > 0; i-- {
			z01.PrintRune(r.side1)
		}
		z01.PrintRune(r._4)
		z01.PrintRune('\n')

	}
}

type rectangle struct {
	_1    rune
	_2    rune
	_3    rune
	_4    rune
	side1 rune
	side2 rune
}
