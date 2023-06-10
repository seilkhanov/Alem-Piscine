package piscine

func JumpOver(str string) string {
	newstring := ""

	if len(str) < 3 {
		return string('\n')
	}

	for i, letter := range str {
		if (i+1)%3 == 0 {
			newstring += string(letter)
		}
	}

	return newstring + "\n"
}
