package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}

	size := max - min
	answer := []int{}

	for i := 0; i < size; i++ {
		answer = append(answer, min)
		min++
	}
	return answer
}
