package piscine

func ConcatParams(args []string) string {
	size := len(args)
	answer := args[0]

	for i := 1; i < size; i++ {
		answer += "\n" + args[i]
	}

	return answer
}
