package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if f(a[j], a[j+1]) < 0 {
				continue
			} else {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
