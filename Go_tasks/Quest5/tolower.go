package piscine

func ToLower(s string) string {
	newword := []byte{}

	for i := range s {
		if s[i] > 64 && s[i] < 91 {
			newword = append(newword, s[i]+32)
		} else {
			newword = append(newword, s[i])
		}
	}
	return string(newword)
}
