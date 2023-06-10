package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, arg := range tab {
		if f(arg) {
			count++
		}
	}
	return count
}
