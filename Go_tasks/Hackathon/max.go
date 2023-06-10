package piscine

func Max(a []int) int {
	number := 0
	for _, arg := range a {
		if arg >= number {
			number = arg
		}
	}
	return number
}
