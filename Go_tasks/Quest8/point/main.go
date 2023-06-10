package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point, number []int) {
	ptr.x = number[0]
	ptr.y = number[1]
}

func main() {
	points := point{}

	number := []int{42, 21}

	setPoint(&points, number)

	output := []rune{'x', ' ', '=', ' ', rune(points.x/10 + '0'), rune(points.x%10 + '0'), ',', ' ', 'y', ' ', '=', ' ', rune(points.y/10 + '0'), rune(points.y%10 + '0'), '\n'}

	for _, r := range output {
		z01.PrintRune(r)
	}
}
