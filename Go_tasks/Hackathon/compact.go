package piscine

func Compact(ptr *[]string) int {
	length := 0
	for _, args := range *ptr {
		if args != "" {
			length++
		}
	}

	newslice := make([]string, length)

	t := 0
	for _, args := range *ptr {
		if args != "" {
			newslice[t] = args
			t++
		}
	}
	*ptr = newslice
	return len(newslice)
}
