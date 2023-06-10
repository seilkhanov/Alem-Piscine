package piscine

func Abort(a, b, c, d, e int) int {
	numbers := []int{a, b, c, d, e}

	SortIntegerTable(numbers)

	return numbers[2]
}

func SortIntegerTable(table []int) {
	for i := 0; i < len(table)-1; i++ {
		for j := 0; j < len(table)-1-i; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}
