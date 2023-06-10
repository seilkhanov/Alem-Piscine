package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	purchases := make(map[string]int, 0)

	list := []string{}
	newword := ""

	if str == "" {
		list = append(list, newword)
	}

	for i, letter := range str {
		if letter != ' ' {
			newword += string(letter)
		} else {
			list = append(list, newword)
			newword = ""
		}

		if i == len(str)-1 {
			list = append(list, newword)
		}
	}

	for _, item := range list {
		purchases[item]++
	}

	return purchases
}
