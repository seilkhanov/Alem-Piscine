package piscine

func Map(f func(int) bool, a []int) []bool {
	output := make([]bool, len(a))
	for i, arg := range a {
		output[i] = f(arg)
	}
	return output
}
