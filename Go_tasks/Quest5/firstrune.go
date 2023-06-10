package piscine

func FirstRune(str string) rune {
	if len(str) == 0 {
		return 0
	}

	word := []rune(str)

	return word[0]
}
