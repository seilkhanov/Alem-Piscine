package piscine

func StringToIntSlice(str string) []int {
	var array []int

	for _, char := range str {
		array = append(array, int(char))
	}
	return array
}
