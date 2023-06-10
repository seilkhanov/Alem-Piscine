package piscine

func NRune(s string, n int) rune {
	if n <= 0 || n > len(s) {
		return rune(0)
	}

	word := []rune(s)

	return word[n-1]
}
