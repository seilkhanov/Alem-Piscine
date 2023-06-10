package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}

	root := 0

	for i := 1; i <= nb; i++ {
		if i*i == nb {
			root = i
			break
		} else if i*i > nb {
			break
		}
	}
	return root
}
