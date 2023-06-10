package piscine

func Join(strs []string, sep string) string {
	joined := ""

	for i := 0; i < len(strs); i++ {
		if i != len(strs)-1 {
			joined += strs[i] + sep
		} else {
			joined += strs[i]
		}
	}
	return joined
}
