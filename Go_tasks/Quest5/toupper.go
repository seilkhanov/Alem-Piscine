package piscine

func ToUpper(s string) string {
	newword := []byte{}

	for i := range s {
		if s[i] > 96 && s[i] < 123 {
			newword = append(newword, s[i]-32)
		} else {
			newword = append(newword, s[i])
		}
	}
	return string(newword)
}
