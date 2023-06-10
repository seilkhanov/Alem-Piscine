package piscine

func Split(s, sep string) []string {
	answer := []string{}
	start := 0

	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			if i-start >= len(sep) {
				answer = append(answer, s[start:i])
			}
			start = i + len(sep)
			i = start - 1
		}
	}

	if start < len(s) {
		answer = append(answer, s[start:])
	}

	return answer
}
