package piscine

func BasicJoin(elems []string) string {
	joined := ""

	for i := range elems {
		joined += elems[i]
	}
	return joined
}
