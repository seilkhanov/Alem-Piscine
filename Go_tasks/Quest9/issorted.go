package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	ascending := 0
	descending := 0
	constant := 0
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) < 0 {
			ascending++
		}

		if f(a[i], a[i+1]) > 0 {
			descending++
		}

		if f(a[i], a[i+1]) == 0 {
			constant++
		}
	}
	if ascending == len(a)-1 || descending == len(a)-1 || constant == len(a)-1 {
		return true
	}
	return false
}

func f(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	}
	return -1
}
