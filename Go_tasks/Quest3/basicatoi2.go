package piscine

func BasicAtoi2(s string) int {
	result := 0
	sign := 1
	i := 0

	if len(s) > 0 && s[0] == '-' {
		sign = -1
		i++
	}

	for ; i < len(s); i++ {
		digit := int(s[i] - '0')

		if digit < 0 || digit > 9 {
			result = 0
			break
		} else {
			result = result*10 + digit
		}
	}
	return result * sign
}
