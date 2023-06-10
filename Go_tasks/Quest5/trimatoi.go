package piscine

func TrimAtoi(s string) int {
	result := 0
	sign := 1
	isdigitfirst := 0

	for i := 0; i < len(s); i++ {
		digit := int(s[i] - '0')

		if digit < 0 || digit > 9 {
			if s[i] == '-' && isdigitfirst == 0 {
				sign = -1
			}
			continue
		} else {
			result = result*10 + digit
			isdigitfirst++
		}
	}
	return result * sign
}
