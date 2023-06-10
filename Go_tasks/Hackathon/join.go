package piscine

func Join(strs []string, sep string) string {
	output := ""
	for i, arg := range strs {
		output += arg
		if i != len(strs)-1 {
			output += sep
		}
	}
	return output
}
