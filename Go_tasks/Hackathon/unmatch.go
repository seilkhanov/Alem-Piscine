package piscine

func Unmatch(a []int) int {
	countMap := make(map[int]int)

	for _, num := range a {
		countMap[num]++
	}

	for num, count := range countMap {
		if count%2 != 0 {
			return num
		}
	}

	return -1
}
