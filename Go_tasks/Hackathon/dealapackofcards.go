package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	for i := 0; i < 4; i++ {
		player := i + 1
		start := i * 3
		end := (i + 1) * 3
		fmt.Printf("Player %d: ", player)
		for j := start; j < end; j++ {
			fmt.Printf("%d", deck[j])
			if j < end-1 {
				fmt.Printf(", ")
			}
		}
		z01.PrintRune('\n')
	}
}
