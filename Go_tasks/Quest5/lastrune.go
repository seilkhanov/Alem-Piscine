package piscine

func LastRune(s string) rune {
	if len(s) == 0 {
		return 0
	}

	word := []rune(s)

	return word[len(s)-1]
}
