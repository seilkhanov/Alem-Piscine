package piscine

func AlphaCount(s string) int {
	count := 0
	for i := range s {
		if (s[i] > 64 && s[i] < 91) || (s[i] > 96 && s[i] < 123) {
			count += 1
		}
	}
	return count
}
