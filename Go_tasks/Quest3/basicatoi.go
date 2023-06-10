package piscine

func BasicAtoi(s string) int {
	result := 0
	sign := 1
	i := 0

	if len(s) > 0 && s[0] == '-' {
		sign = -1
		i++
	}

	for ; i < len(s); i++ {
		digit := int(s[i] - '0')
		result = result*10 + digit
	}
	return result * sign
}
