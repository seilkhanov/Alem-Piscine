package piscine

func Compare(a, b string) int {
	minlength := len(a)

	if len(a) > len(b) {
		minlength = len(b)
	}

	for i := 0; i < minlength; i++ {
		if a[i] > b[i] {
			return 1
		} else if a[i] < b[i] {
			return -1
		}
	}

	if len(a) > len(b) {
		return 1
	} else if len(a) < len(b) {
		return -1
	}

	return 0
}
