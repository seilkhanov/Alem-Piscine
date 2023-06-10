package piscine

func StrRev(s string) string {
	one := []rune(s)
	reverse := make([]rune, 0)
	for i := len(one) - 1; i > -1; i-- {
		reverse = append(reverse, one[i])
	}
	return string(reverse)
}
