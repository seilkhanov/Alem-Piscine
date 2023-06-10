package piscine

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}

	if nb == 2 || nb == 3 {
		return true
	}

	if nb%2 == 0 || nb%3 == 0 {
		return false
	}

	divisor := 5

	for divisor*divisor <= nb {
		if nb%divisor == 0 || nb%(divisor+2) == 0 {
			return false
		}
		divisor += 6
	}

	return true
}
