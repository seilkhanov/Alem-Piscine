package piscine

func DescendAppendRange(max, min int) []int {
	list := []int{}
	if max <= min {
		return list
	} else {
		for i := max; i > min; i-- {
			list = append(list, i)
		}
	}
	return list
}
