package piscine

func ActiveBits(n int) int {
	count := 1
	for n > 1 {
		if n%2 == 1 {
			count++
		}

		n = n / 2
	}
	return count
}
