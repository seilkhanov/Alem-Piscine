package piscine

func Any(f func(string) bool, a []string) bool {
	for _, arg := range a {
		if f(arg) {
			return true
		}
	}
	return false
}
